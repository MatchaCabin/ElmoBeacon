import json
import os
import sys 
from mitmproxy import ctx, http
from mitmproxy.tools.main import mitmdump 

# --- Configuration ---
# For now, targeting US specifically as it's the one confirmed affected.
TARGET_HOSTS = [
    "gf2-gacha-record-us.sunborngame.com",
    # Add other oversea gacha domains here if they also become affected
    # "gf2-gacha-record-intl.sunborngame.com",
    # "gf2-gacha-record-jp.sunborngame.com",
]
TARGET_PATH_PREFIX = "/list" # The common path for the gacha list endpoint
OUTPUT_FILE = "gacha_session_info.json"

# To prevent multiple writes in a single mitmproxy session.
#info_saved_this_session = False

class GachaInfoSaver:
    def request(self, flow: http.HTTPFlow) -> None:
        global info_saved_this_session
        #if info_saved_this_session:
        #    return

        req = flow.request

        if req.host in TARGET_HOSTS and \
           req.path.startswith(TARGET_PATH_PREFIX) and \
           req.method == "POST":

            try:
                # Looking for the initial request for a banner type
                form_data = req.urlencoded_form
                if "next" not in form_data and "type_id" in form_data:
                    ctx.log.info(f"Potential initial gacha request detected: {req.pretty_url}")

                    access_token = req.headers.get("Authorization", None)
                    # req.pretty_url gives the full URL including scheme, host, path, and query parameters
                    full_gacha_url = req.pretty_url

                    if access_token and full_gacha_url:
                        data_to_save = {
                            "FullGachaRecordUrl": full_gacha_url,
                            "AccessToken": access_token,
                            "UidString": "REPLACE_WITH_YOUR_UID"
                        }

                        # Save the file in the same directory as the script for simplicity
                        script_dir = os.path.dirname(os.path.abspath(__file__))
                        output_path = os.path.join(script_dir, OUTPUT_FILE)

                        with open(output_path, "w") as f:
                            json.dump(data_to_save, f, indent=4)
                        
                        print("-" * 70)
                        print(f"[SUCCESS] Gacha session info saved to: {output_path}")
                        print(f"  FullGachaRecordUrl: {full_gacha_url}")
                        print(f"  AccessToken: {access_token[:20]}... (truncated for display)")
                        print(f"  UidString: PLEASE OPEN '{OUTPUT_FILE}' AND REPLACE 'REPLACE_WITH_YOUR_UID' with your actual game UID.")
                        print("-" * 70)
                        print("[IMPORTANT]")
                        print("you MUST turn OFF the system proxy settings now that the script is finished")
                        print("or your internet connection will NOT work correctly")
                        print("(Windows: Settings > Network & Internet > Manual proxy)")
                        print("-" * 70)
                        print("mitmproxy will attempt to shut down now...")
                        #info_saved_this_session = True # Prevent further writes in this session
                        ctx.master.shutdown()
                    else:
                        if not access_token:
                            ctx.log.warn("Authorization header missing in the targeted gacha request.")
                        # full_gacha_url should always be present if req.pretty_url works
            except Exception as e:
                # Log any errors during form data processing or file writing
                ctx.log.error(f"Error processing request or saving gacha info: {e}")

addons = [GachaInfoSaver()]

def main():
    print("-" * 50)
    print("Starting mitmdump with GachaInfoSaver addon...")
    print(f"Ensure your system/game proxy is set to http://127.0.0.1:8080")
    print(f"And the mitmproxy CA certificate is installed and trusted.")
    print(f"Trigger the gacha history view in your game for the US server.")
    print(f"The script will capture the details and save them to '{OUTPUT_FILE}'.")
    print("Press Ctrl+C to stop mitmdump if it doesn't shut down automatically.")
    print("-" * 50)
    
    # Can add filters to the command line here.
    mitm_args = ['-s', __file__]
    
    try:
        mitmdump(mitm_args)
    except KeyboardInterrupt:
        print("\nmitmdump stopped by user.")
    sys.exit(0)

if __name__ == "__main__":
    # This allows the script to be run directly using `python save_gacha_info.py`
    main()
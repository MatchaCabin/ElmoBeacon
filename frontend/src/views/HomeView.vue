<script setup lang="ts">
</script>

<template>
  <div class="flex flex-col gap-4">
    <section>
      <h2 class="text-2xl font-bold">{{$t('home.intro.title')}}</h2>
      <div class="text-md">
        <p><Link href="https://gf2.mcc.wiki">ElmoBeacon</Link> is a tool for storing and analyzing GFL2 gacha records.</p>
        <p>The project is compatible with DarkWinter server(China,North America) and Haoplay server(Global,Asia,Japan,Korea).</p>
        <p>This is an open source project and maintained by <Link href="https://gf2.mcc.wiki">MccWiki</Link> and the community.</p>
        <p>The program has implemented the basic features. UI will be optimized later.</p>
        <p>If you have any ideas for UI, feel free to PR!</p>
      </div>
    </section>

    <section>
      <h2 class="text-2xl font-bold">{{$t('home.requirement.title')}}</h2>
      <div class="text-md">
        <p>1.<Link href="https://developer.microsoft.com/en-us/microsoft-edge/webview2">WebView2</Link>(embedded)</p>
        <p>2.GFL2 Client</p>
      </div>
    </section>

    <section>
      <h2 class="text-2xl font-bold">{{$t('home.instruction.title')}}</h2>
      <div class="text-md">
        <p>1.Ensure that you have had a recent valid game login on this device.</p>
        <p>2.Switch to the record page in the left menu bar.</p>
        <p>3.Click incremental update or full update depending on your needs.</p>
      </div>
    </section>
  </div>
</template>
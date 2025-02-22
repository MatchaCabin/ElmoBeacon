export namespace handler {
	
	export class DisplayRecord {
	    Name: string;
	    Icon: string;
	    Count: number;
	    IsMissing: boolean;
	
	    static createFrom(source: any = {}) {
	        return new DisplayRecord(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Icon = source["Icon"];
	        this.Count = source["Count"];
	        this.IsMissing = source["IsMissing"];
	    }
	}
	export class PoolInfo {
	    storedCount: number;
	    recordList: DisplayRecord[];
	    totalCount: number;
	    rank5Count: number;
	    rank4Count: number;
	    rank3Count: number;
	    rank5Rate: number;
	    rank4Rate: number;
	    rank3Rate: number;
	    rank5Avg: number;
	    rank5UpAvg: number;
	    missingCount: number;
	    missingRate: number;
	
	    static createFrom(source: any = {}) {
	        return new PoolInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.storedCount = source["storedCount"];
	        this.recordList = this.convertValues(source["recordList"], DisplayRecord);
	        this.totalCount = source["totalCount"];
	        this.rank5Count = source["rank5Count"];
	        this.rank4Count = source["rank4Count"];
	        this.rank3Count = source["rank3Count"];
	        this.rank5Rate = source["rank5Rate"];
	        this.rank4Rate = source["rank4Rate"];
	        this.rank3Rate = source["rank3Rate"];
	        this.rank5Avg = source["rank5Avg"];
	        this.rank5UpAvg = source["rank5UpAvg"];
	        this.missingCount = source["missingCount"];
	        this.missingRate = source["missingRate"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SyncResult {
	    Server: string;
	    Uid: number;
	    PoolType: number;
	    Count: number;
	
	    static createFrom(source: any = {}) {
	        return new SyncResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Server = source["Server"];
	        this.Uid = source["Uid"];
	        this.PoolType = source["PoolType"];
	        this.Count = source["Count"];
	    }
	}

}

export namespace model {
	
	export class User {
	    id: number;
	    uid: number;
	    server: string;
	    gameDataDir: string;
	    lastBBSToken: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.uid = source["uid"];
	        this.server = source["server"];
	        this.gameDataDir = source["gameDataDir"];
	        this.lastBBSToken = source["lastBBSToken"];
	    }
	}

}


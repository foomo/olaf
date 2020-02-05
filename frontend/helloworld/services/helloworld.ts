/* tslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package

export class ServiceClient {
	public static defaultEndpoint = "/services/helloworld";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async helloWorld(name:string):Promise<string> {
		return (await this.transport<{0:string}>("HelloWorld", [name]))[0]
	}
}
export class AdminServiceClient {
	public static defaultEndpoint = "/services/helloworld-admin";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async helloAdmin():Promise<string> {
		return (await this.transport<{0:string}>("HelloAdmin", []))[0]
	}
}
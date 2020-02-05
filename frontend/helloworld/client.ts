import axios, { AxiosResponse } from "axios";
// axios transport
let getTransport = endpoint => async <T>(method, args = []) => {
  return new Promise<T>(async (resolve, reject) => {
    try {
      let axiosPromise: AxiosResponse<T> = await axios.post<T>(
        endpoint + "/" + encodeURIComponent(method),
        JSON.stringify(args)
      );
      return resolve(axiosPromise.data);
    } catch (e) {
      return reject(e);
    }
  });
};

type Transport = ReturnType<typeof getTransport>;

interface ClientConstructor<T> {
	defaultEndpoint:string;
    new (transport:Transport):T;
}

export const getClient = (serverProvider:()=>string) => <T>(ClientClass:ClientConstructor<T>) =>
  new ClientClass(
    getTransport(serverProvider() + ClientClass.defaultEndpoint)
  );

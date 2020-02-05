import { useState } from "react";
import {getClient} from "../client";
import { ServiceClient, AdminServiceClient } from "../services/helloworld";

const serverProvider = () => {
  if("undefined" === typeof window) {
    return "http://localhost:8080";
  } else {
    return "http://localhost:8080"; // ""
  }
}

const client = getClient(serverProvider)(ServiceClient);
const adminClient = getClient(serverProvider)(AdminServiceClient);

const defaultGreeting = () => {
  if("undefined" === typeof window) {
    return "I am server";
  } else {
    return "me client"; // ""
  }
}

const Home = () => {
  const [name, setName] = useState("anonymous");
  const [greeting, setGreeting] = useState(defaultGreeting());
  return (
    <>
      <h1>Hello world! {name}</h1>
      <div>
        <input value={name} onChange={e => {
            setName(e.target.value);
            client.helloWorld(e.target.value).then(reply => setGreeting(reply));
            adminClient.helloAdmin().then(hello => console.log("admin said", hello));
        }}></input>
        <div>
          <h2>Greetings from the service</h2>
          <pre>{greeting}</pre>
        </div>
      </div>
    </>
  );
};

export default Home;

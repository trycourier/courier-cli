import store from "./configStore.js";
import axios from "axios";

const randomJoke = async () => {
  var options = {
    method: 'GET',
    url: 'https://v2.jokeapi.dev/joke/Any?safe-mode',
    params: {
      format: 'json',
      contains: 'C%23',
      idRange: '0-150',
      blacklistFlags: 'nsfw,racist'
    },
    headers: {
      'x-rapidapi-host': 'jokeapi-v2.p.rapidapi.com',
    }
  };
  
  const response = await axios.request(options);
  let joke = ""
  if(response.data.type == "twopart"){
      joke = `${response.data.setup} ${response.data.delivery}`;
  }else{
      joke = String(response.data.joke) + String(response.data.setup)
  }
  store.set('randomJoke', joke);

  return joke;
}

export default randomJoke;
<script>
  import Home from "./pages/Home.svelte"

  export let path_api = "";
  export let version = "";

  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  const token_browser = urlParams.get("token");
  
  if (token_browser === null) {
    console.log("TOKEN NOT FOUND");
  } else {
    initTimezone();
  }

  let flag_game = true;
  let client_credit = 0;
  let client_ipaddress = "0.0.0.0";
  let client_timezone = "";
  let client_company = "";
  let client_username = "";
  let client_name = "";
  let engine_time = 0;
  let engine_invoice = "";
  let engine_status = "LOCK";

  async function initTimezone() {
    const res = await fetch(path_api+"api/healthz");
    if (!res.ok) {
      const message = `An error has occured: ${res.status}`;
      throw new Error(message);
    } else {
      const json = await res.json();
      client_ipaddress = json.real_ip;
      client_timezone = "Asia/Jakarta";
    }
    // initapp(token_browser);
  }
  async function initapp(token) {
      const res = await fetch(path_api+"api/checktoken", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            client_token: token,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
          // logout();
      } else if (json.status == 403) {
          // alert(json.message);
      } else {
          flag_game = true;
          client_company = json.client_company;
          client_name = json.client_name;
          client_username = json.client_username;
          client_credit = json.client_credit;
      }
  }
  var source = new EventSource(path_api+"sse");
  source.onmessage = function(event) {
    let text_dasar = event.data;
    let text_result = text_dasar.split(":");
    let text_final = text_result[1].replace(`"`,"")
    let text_finalreplace = text_final.replace(`"`,"")
    let text_finalsplit = text_finalreplace.split("|");
    // let text_finalsplit2 = text_finalsplit.split("|");
    console.log(text_finalsplit)
    let data_invoice = text_finalsplit[0].replace(` `,"");
    let data_time = text_finalsplit[1];
    let data_status = text_finalsplit[3];
  
    engine_time = data_time;
    engine_invoice = data_invoice;
    engine_status = data_status;
    
  };
  
</script>
<main class="container mx-auto px-2 mt-5 text-base-content   xl:mt-7 max-w-screen-xl  pb-5 h-fit lg:h-full">

  {#if flag_game}
  <Home 
    {engine_time}  
    {engine_invoice}  
    {engine_status}  />
  {/if}
</main>

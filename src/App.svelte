<script>
  import dayjs from "dayjs";
  import utc from "dayjs/plugin/utc";
  import timezone from "dayjs/plugin/timezone";

  
  dayjs.extend(utc);
  dayjs.extend(timezone);

  let client_timezone = "Asia/Jakarta"
  let client_name = "Developer";
  let client_ipaddress = "127.127.127.127";
  let client_credit = 10000;
  let engine_invoice = "202402272865"
  let engine_time = 30
  let engine_minbet = 500
  let engine_maxbet = 2000000
  let clockmachine = "";

  let field_maxlength_bet = length
  let field_nomor = ""



  function updateClock() {
    let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");
    clockmachine = endtime;
    
  }

  var x = setInterval(function() {
    engine_time = engine_time - 1;
    if (engine_time < 0) {
      clearInterval(x);
    }
  }, 1000);

  $: {
        setInterval(updateClock, 1000);
       
  }
</script>
<main class="container mx-auto px-2 mt-5 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 h-fit lg:h-full">
  <section class="flex-col w-full p-2 rounded-md ">
    <img class="w-[150px]" src="https://i.imgur.com/PNSe1ov.png" alt="" srcset="">
    <section class="flex justify-between w-full">
      <h1 class="text-2xl font-bold ">
        #{engine_invoice} <br />
        {engine_time} S 
      </h1>
      <div>
        <p class="w-full text-xs lg:text-sm text-right select-none">
          Asia/Jakarta <br />
          {clockmachine}  WIB (+7)<br>
          {client_name} <br />
          {client_ipaddress}
        </p>
        <div class="w-full text-xs lg:text-sm text-right select-none">
          CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(client_credit)}</span>
        </div>
      </div>
    </section>
    <section class="grid w-full gap-2 mt-10">
      <div class="flex w-full bg-base-300">
        <input 
          class="w-full text-[50px] text-center link-accent bg-base-300 focus:outline-none"
          type="text" 
          placeholder="2D" 
          maxlength="2"
          value="{field_nomor}">
      </div>
      <div class="flex w-full bg-base-300">
        <span class="pt-3 ml-2 text-green-500">Rp</span>
        <input 
          class="w-1/2 h-10 p-2 text-xl ml-1 bg-base-300 link-accent focus:outline-none"
          type="text" 
          placeholder="Bet" 
          maxlength="7"
          value="{engine_minbet}">
          <div class="flex gap-2 w-full justify-end py-1 mr-2">
            <button class="btn btn-sm btn-active">Min</button>
            <button class="btn btn-sm btn-active">Max</button>
            <button class="btn btn-sm btn-active">1/2</button>
            <button class="btn btn-sm btn-active">x2</button>
          </div>
      </div>
      
      <button class="btn btn-success">Bayar</button>
    </section>
    
    
  </section>
</main>
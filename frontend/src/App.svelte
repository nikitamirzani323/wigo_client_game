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
  let field_bet = engine_minbet
  let field_nomor = ""

  function updateClock() {
    let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");
    clockmachine = endtime;
    
  }

  var x = setInterval(function() {
    engine_time = engine_time - 1;
    if (engine_time < 1) {
      clearInterval(x);
    }
  }, 1000);

  $: {
        setInterval(updateClock, 1000);
       
  }
  const call_buttonbet = (e) => {
    switch(e){
      case "min":
        field_bet = engine_minbet
        break;
      case "max":
        field_bet = engine_maxbet
        break;
      case "1/2":
        field_bet = engine_minbet + (parseInt(engine_minbet)*0.5);
        break;
      case "2":
        field_bet = engine_minbet + (parseInt(engine_minbet)*2);
        break;
    }
  };
  const handleKeyboard_number = (e) => {
    if (isNaN(parseInt(e.target.value))) {
        return e.target.value = "";
    }else{
        return e.target.value = parseInt(e.target.value);
    }
  }
</script>
<main class="container mx-auto px-2 mt-5 text-base-content   xl:mt-7 max-w-screen-xl  pb-5 h-fit lg:h-full">
  <section class="glass bg-opacity-60 xl:rounded-box">
    <section class="flex-col w-full p-2 rounded-md ">
      <img class="w-[150px]" src="https://i.imgur.com/PNSe1ov.png" alt="" srcset="">
      <section class="flex justify-between w-full">
        <section class="flex-col  font-bold  w-1/2 rounded-md select-none">
          <center>
            <div class="text-2xl">
              #{engine_invoice} 
            </div>
            <div class="text-[50px] pt-1">
              {engine_time} S 
            </div>
          </center>
        </section>
        <section class="w-full">
          <p class="w-full text-xs lg:text-sm text-right select-none">
            Asia/Jakarta <br />
            {clockmachine}  WIB (+7)<br>
            {client_name} <br />
            {client_ipaddress}
          </p>
          <div class="w-full text-xs lg:text-sm text-right select-none">
            CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(client_credit)}</span>
          </div>
        </section>
      </section>
      <section class="grid w-full gap-2 mt-10">
        <div class="flex w-full bg-base-300">
          <input on:keyup={handleKeyboard_number}
            bind:value={field_nomor}  
            class="w-full text-[50px] text-center link-accent bg-base-300 focus:outline-none"
            type="text" 
            placeholder="2D" 
            maxlength="2">
        </div>
        <div class="flex-col">
          <div class="flex w-full bg-base-300">
            <span class="pt-2 ml-2 text-xl text-green-500">Rp</span>
            <input on:keyup={handleKeyboard_number}
              bind:value={field_bet}  
              class="w-1/2 h-10 p-2 text-2xl ml-1 bg-base-300 link-accent focus:outline-none"
              type="text" 
              placeholder="Bet" 
              maxlength="7">
              <div class="flex gap-2 w-full justify-end py-1 mr-2">
                <button on:click={() => {
                      call_buttonbet("min");
                    }}
                  class="btn btn-sm btn-active">Min</button>
                <button on:click={() => {
                    call_buttonbet("max");
                  }}
                  class="btn btn-sm btn-active">Max</button>
                <button on:click={() => {
                    call_buttonbet("1/2");
                  }} class="btn btn-sm btn-active">1/2</button>
                <button on:click={() => {
                    call_buttonbet("2");
                }} class="btn btn-sm btn-active">x2</button>
              </div>
          </div>
          <span class="pl-11 -pt-1 text-[11px] link-accent">{new Intl.NumberFormat().format(field_bet)}</span>
        </div>
        
        
        <button class="btn btn-success">
          Bayar 
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
          </svg>
          
        </button>
      </section>
    </section>
  </section>
  <section class="flex gap-2 mt-4">
    <button class="btn ">Taruhan Saya</button>
    <button class="btn ">Riwayat</button>
  </section>
</main>

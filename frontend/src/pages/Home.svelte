<script>
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";
  
    
    dayjs.extend(utc);
    dayjs.extend(timezone);
   
    export let path_api = "";
    export let version = "";
    export let engine_time = 0
    export let engine_invoice = ""
    export let engine_status = "LOCK"
    
    let flag_toast = false;
    let toast_message = "";
    let client_company = "nuke"
    let client_username = "developer"
    let client_timezone = "Asia/Jakarta"
    let client_name = "Developer";
    let client_ipaddress = "127.127.127.127";
    let client_credit = 10000;
    
    let engine_minbet = 500
    let engine_maxbet = 2000000
    let engine_multiplier = 5.2
    let clockmachine = "";
    
    let flag_btnbuy = false;
    let field_maxlength_bet = length
    let field_bet = engine_minbet
    let field_nomor = ""
    
    let list_invoice = []
    let list_result = []
    let flag_listinvoice = true;
    let flag_listresult = false;

    function updateClock() {
      let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");
      clockmachine = endtime;
      
    }
  
    function toast_hidden() {
        flag_toast = false;
    }

    
  
    $: {
        setInterval(updateClock, 1000);
        fetch_invoiceall()
        
    }
    async function call_bayar() {
        let flag = true;
        let msg_err = ""
        let prize2d = field_nomor.length;
        if(field_nomor == ""){
            flag = false
            msg_err = "Nomor wajib diisi"
        }
        if (prize2d.toString() == "1") {
            flag = false
            msg_err = "Nomor minimal 2 digit"
        }
        if(field_bet == ""){
            flag = false
            msg_err = "Bet wajib diisi"
        }
        if(parseInt(field_bet) < parseInt(engine_minbet)){
            flag = false
            msg_err = "Minimal Bet " + engine_minbet
        }
        if(parseInt(field_bet) > parseInt(engine_maxbet)){
            flag = false
            msg_err = "Maximal Bet " + engine_maxbet
        }
        if(parseInt(field_bet) > parseInt(client_credit)){
            flag = false
            msg_err = "Credit tidak cukup "
        }
        if(flag){
            const res = await fetch(path_api+"api/savetransaksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    transaksidetail_company: client_company,
                    transaksidetail_idtransaksi: engine_invoice,
                    transaksidetail_username: client_username,
                    transaksidetail_nomor: field_nomor,
                    transaksidetail_bet: parseInt(field_bet),
                    transaksidetail_multiplier: parseFloat(engine_multiplier),
                }),
            });
            const json = await res.json();
            if (json.status === 400) {
                // logout();
            } else if (json.status == 403) {
                alert(json.message);
            } else {
                client_credit = parseInt(client_credit) - parseInt(field_bet)
                field_bet = engine_minbet
                field_nomor = ""
                fetch_invoiceall()
                flag_toast = true
                toast_message = json.message
            }
        }else{
            flag_toast = true
            toast_message = msg_err
        }
        setTimeout(toast_hidden, 3000);
    }
    async function fetch_listresult() {
        flag_listinvoice = false;
        flag_listresult = true;
        list_result = []
        const res = await fetch(path_api+"api/listresult", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                invoice_company: client_company,
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
        
        } else if (json.status == 403) {
            alert(json.message);
        } else {
        let record = json.record;
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
                list_result = [
                    ...list_result,
                    {
                        result_invoice: record[i]["result_invoice"],
                        result_date: record[i]["result_date"],
                        result_result: record[i]["result_result"],
                    },
                ];
            }
        }
        }
    }
    async function fetch_invoiceall() {
        flag_listinvoice = true;
        flag_listresult = false;
        list_invoice = []
        const res = await fetch(path_api+"api/listinvoice", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                invoice_company: client_company,
                Invoice_username: client_username,
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
        
        } else if (json.status == 403) {
            alert(json.message);
        } else {
        let record = json.record;
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
                let status_css = ""
                
                switch(record[i]["invoiceclient_status"]){
                    case "LOSE":
                        status_css = "bg-primary text-white";
                        break;
                    case "WIN":
                        status_css = "bg-success text-white";
                        break;
                    case "RUNNING":
                        status_css = "bg-accent text-white";
                        break;
                }

               
                list_invoice = [
                    ...list_invoice,
                    {
                        invoiceclient_id: record[i]["invoiceclient_id"],
                        invoiceclient_date: record[i]["invoiceclient_date"],
                        invoiceclient_result: record[i]["invoiceclient_result"],
                        invoiceclient_nomor: record[i]["invoiceclient_nomor"],
                        invoiceclient_bet: record[i]["invoiceclient_bet"],
                        invoiceclient_multiplier: record[i]["invoiceclient_multiplier"],
                        invoiceclient_win: record[i]["invoiceclient_win"],
                        invoiceclient_status: record[i]["invoiceclient_status"],
                        invoiceclient_status_css: status_css,
                    },
                ];
            }
        }
        }
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
    const call_allinvoice = () => {
        fetch_invoiceall()
    };
    const call_listresult = () => {
        fetch_listresult()
    };
    const handleKeyboard_nomor = (e) => {
      if (isNaN(parseInt(e.target.value))) {
          return e.target.value = "";
      }else{
          return e.target.value = e.target.value;
      }
    }
    const handleKeyboard_number = (e) => {
      if (isNaN(parseInt(e.target.value))) {
          return e.target.value = "";
      }else{
          return e.target.value = parseInt(e.target.value);
      }
    }
    const handleKeyboard_number_blur = (e) => {
        let temp_number = e.target.value
        if (isNaN(parseInt(e.target.value))) {
            return e.target.value = "";
        }else{
            if(parseInt(temp_number) < parseInt(engine_minbet)){
                return e.target.value = parseInt(engine_minbet);
            }else{
                if(parseInt(temp_number) > parseInt(engine_maxbet)){
                    return e.target.value = parseInt(engine_maxbet);
                }
            }
            
        }
    }
    
  </script>
 
<section class="glass bg-opacity-60 rounded-lg">
    <section class="flex-col w-full p-2 rounded-md ">
        <center>
            <img class="w-[150px]" src="https://i.imgur.com/PNSe1ov.png" alt="" srcset="">
        </center>
        <section class="hidden lg:flex justify-between w-full bg-base-100 p-1 rounded-md select-none mt-1">
            <section class="flex-col text-center  font-bold  w-1/2  ">
                <div class="flex-col">
                    <div class="text-lg">PERIODE</div>
                    <div class="link-accent text-sm">{engine_invoice}</div>
                </div>
                <div class="flex-col mt-2">
                    <div class="text-lg">WAKTU</div>
                    <div class="link-accent text-sm">{engine_time} S </div>
                </div>
            </section>
            <section class="w-full ">
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
        <section class="flex-col lg:hidden w-full ">
            <section class="flex justify-between  w-full  bg-base-100 p-2 rounded-md select-none mt-1">
                <div class="flex-col w-full text-center">
                    <div class="text-sm">PERIODE</div>
                    <div class="link-accent text-xs">{engine_invoice}</div>
                </div>
                <div class="flex-col w-full text-center">
                    <div class="text-sm">WAKTU</div>
                    <div class="link-accent text-xs">{engine_time} S </div>
                </div>
            </section>
            <section class="flex w-full bg-base-100 p-2 rounded-md select-none mt-1">
                <p class="w-full text-xs lg:text-sm text-left select-none">
                    Asia/Jakarta <br />
                    {clockmachine}  WIB (+7)<br>
                    {client_name} <br />
                    {client_ipaddress}<br />
                    CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(client_credit)}</span>
                </p>
            </section>
        </section>
        <section class="grid w-full gap-2 mt-2">
            <div class="flex w-full bg-base-300">
            <input on:keyup={handleKeyboard_nomor}
                bind:value={field_nomor}  
                class="w-full text-[50px] text-center link-accent bg-base-300 focus:outline-none"
                type="text" 
                placeholder="2D" 
                maxlength="2">
            </div>
            <div class="flex-col">
            <div class="flex w-full bg-base-300">
                <span class="pt-2 ml-2 text-xl text-green-500">Rp</span>
                <input 
                    on:keyup={handleKeyboard_number}
                    on:blur={handleKeyboard_number_blur}
                    bind:value={field_bet}  
                    class="w-1/2 h-10 p-2 text-2xl ml-1 bg-base-300 link-accent focus:outline-none"
                    type="text" 
                    placeholder="Bet" 
                    maxlength="7">
                <div class="hidden lg:flex gap-2 w-full justify-end py-1 mr-2">
                    <button on:click={() => {
                        call_buttonbet("min");
                        }}
                    class="btn btn-sm btn-active btn-sm">Min</button>
                    <button on:click={() => {
                        call_buttonbet("max");
                    }}
                    class="btn btn-sm btn-active btn-sm">Max</button>
                    <button on:click={() => {
                        call_buttonbet("1/2");
                    }} class="btn btn-sm btn-active btn-sm">1/2</button>
                    <button on:click={() => {
                        call_buttonbet("2");
                    }} class="btn btn-sm btn-active btn-sm">x2</button>
                </div>
            </div>
            <span class="pl-11 -pt-1 text-[11px] link-accent">{new Intl.NumberFormat().format(field_bet)}</span>
            </div>
            
            {#if engine_status == "OPEN"}
            <button on:click={() => {
                        call_bayar();
                }}  class="btn btn-success">
                Bayar 
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
                </svg>
            </button>
            {/if}
        </section>
    </section>
</section>
<section class="flex-col gap-2 mt-4 p-2 glass bg-opacity-60 rounded-md">
    <div class="flex gap-2">
        <button on:click={() => {
            call_allinvoice();
         }}  class="btn btn-sm">Taruhan Saya</button>
        <button on:click={() => {
            call_listresult();
         }}  class="btn btn-sm">Riwayat</button>
    </div>
    <section class="  mt-4 p-1">
        {#if flag_listinvoice}
        <div class="overflow-auto h-[500px]  scrollbar-thin scrollbar-thumb-green-100">
            <table class="table table-xs  w-full " >
                <thead class="sticky top-0">
                    <tr class="border-none">
                        <th width="5%" class="text-xs text-center align-top">STATUS</th>
                        <th width="5%" class="text-xs text-left align-top">INVOICE</th>
                        <th width="5%" class="text-xs text-center align-top">DATE</th>
                        <th width="10%" class="text-xs text-center align-top">RESULT</th>
                        <th width="*" class="text-xs text-center align-top">NOMOR</th>
                        <th width="10%" class="text-xs text-right align-top">BET</th>
                        <th width="10%" class="text-xs text-right align-top">WIN</th>
                    </tr>
                </thead>
                <tbody>
                    {#each list_invoice as rec}
                    <tr class="border-none">
                        <td class="text-xs  text-center whitespace-nowrap align-top">
                            <span class="{rec.invoiceclient_status_css} p-1 text-xs lg:text-sm  uppercase  rounded-lg w-20 ">{rec.invoiceclient_status}</span>
                        </td>
                        <td class="text-xs  text-left whitespace-nowrap align-top border-b-transparent">{rec.invoiceclient_id}</td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">{rec.invoiceclient_date}</td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">{rec.invoiceclient_result}</td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">{rec.invoiceclient_nomor}</td>
                        <td class="text-xs text-right  whitespace-nowrap align-top link-accent {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoiceclient_bet)}</td>
                        <td class="text-xs text-right  whitespace-nowrap align-top link-secondary {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoiceclient_win)}</td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        {/if}
        {#if flag_listresult}
        <div class="overflow-auto h-[500px]  ">
            <table class="table table-xs w-full " >
                <thead class="sticky top-0">
                    <tr class="border-none">
                        <th width="5%" class="text-xs text-left align-top">INVOICE</th>
                        <th width="5%" class="text-xs text-center align-top">DATE</th>
                        <th width="10%" class="text-xs text-center align-top">RESULT</th>
                    </tr>
                </thead>
                <tbody>
                    {#each list_result as rec}
                    <tr class="border-none">
                        <td class="text-xs  text-left whitespace-nowrap align-top border-b-transparent">{rec.result_invoice}</td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">{rec.result_date}</td>
                        <td class="text-xs  text-center whitespace-nowrap align-top">{rec.result_result}</td>
                    </tr>
                    {/each}
                </tbody>
            </table>
        </div>
        {/if}
    </section>
</section>
{#if flag_toast}
    <div class="toast toast-top toast-center">
        <div class="alert ">
            <span>{toast_message}</span>
        </div>
    </div>
{/if}
  
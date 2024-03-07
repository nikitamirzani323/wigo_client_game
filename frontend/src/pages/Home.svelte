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
    
    let flag_btnbuy = true;
    let field_maxlength_bet = length
    let field_bet = engine_minbet
    let field_nomor = ""
    
    let list_invoice = []
    let list_result = []
    let keranjang = [];
    let flag_listinvoice = true;
    let flag_listresult = false;
    let bet_multiple = ""
    let isModalMinBet = false;
    let isModalInfo = false;
    let client_listbet = [
        {id:500,val:"500"},
        {id:1000,val:"1000"},
        {id:1500,val:"1500"},
        {id:2000,val:"2000"},
        {id:2500,val:"2500"},
        {id:3000,val:"3000"},
        {id:3500,val:"3500"},
        {id:4000,val:"4000"},
        {id:4500,val:"4500"},
        {id:5000,val:"5000"},
        {id:5500,val:"5500"},
        {id:6000,val:"6000"},
        {id:6500,val:"6500"},
        {id:7000,val:"7000"},
        {id:7500,val:"7500"},
        {id:8000,val:"8000"},
        {id:8500,val:"8500"},
        {id:9000,val:"9000"},
        {id:9500,val:"9500"},
        {id:10000,val:"10000"},
        {id:15000,val:"15000"},
        {id:20000,val:"20000"},
        {id:25000,val:"25000"},
        {id:30000,val:"30000"},
        {id:35000,val:"35000"},
        {id:40000,val:"40000"},
        {id:45000,val:"45000"},
        {id:50000,val:"50000"},
        {id:55000,val:"55000"},
        {id:60000,val:"60000"},
        {id:65000,val:"65000"},
        {id:70000,val:"70000"},
        {id:75000,val:"75000"},
        {id:80000,val:"80000"},
        {id:85000,val:"85000"},
        {id:90000,val:"90000"},
        {id:95000,val:"95000"},
        {id:100000,val:"100000"},
        {id:200000,val:"200000"},
        {id:300000,val:"300000"},
        {id:400000,val:"400000"},
        {id:500000,val:"500000"},
        {id:600000,val:"600000"},
        {id:700000,val:"700000"},
        {id:800000,val:"800000"},
        {id:900000,val:"900000"},
        {id:1000000,val:"1000000"},
        {id:1500000,val:"1500000"},
        {id:2000000,val:"2000000"},
        {id:2500000,val:"2500000"},
    ]

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
        // let mergednomor = [...bet_multiple, ...bet_multiple2, ...bet_multiple3, ...bet_multiple4]
        let total_bet_multiple = bet_multiple.length
        let total_bayar = parseInt(total_bet_multiple)*parseInt(field_bet)
        // console.log(bet_multiple.toString())
        // console.log("total nomor : ",total_bet_multiple)
        // console.log("total bayar : ",total_bayar)

        if(parseInt(engine_time) < 5){
            flag = false
            msg_err = "Timeout"
        }
        
        if(field_bet == ""){
            flag = false
            msg_err = "Bet wajib diisi"
        }
        if(parseInt(total_bayar) <= 0){
            flag = false
            msg_err = "Nomor dan Bet wajib diisi"
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
        if(parseInt(total_bayar) > parseInt(client_credit)){
            flag = false
            msg_err = "Credit tidak cukup "
        }
        // flag = false;
        if(flag){
            flag_btnbuy = false;
            for(let i=0;i<total_bet_multiple;i++){
                const data = {
                    nomor: bet_multiple[i],
                    bet: parseInt(field_bet),
                    multiplier: parseFloat(engine_multiplier)
                };
                keranjang = [data, ...keranjang];
            }
            console.log(keranjang)

            const res = await fetch(path_api+"api/savetransaksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    transaksidetail_company: client_company,
                    transaksidetail_idtransaksi: engine_invoice,
                    transaksidetail_username: client_username,
                    transaksidetail_totalbet: parseInt(total_bayar),
                    transaksidetail_listdatabet: keranjang,
                }),
            });
            const json = await res.json();
            if (json.status === 400) {
                flag_btnbuy = true;
            } else if (json.status == 403) {
                alert(json.message);
                flag_btnbuy = true;
            } else {
                client_credit = parseInt(client_credit) - parseInt(total_bayar)
                field_bet = engine_minbet
                field_nomor = ""
                fetch_invoiceall()
                flag_toast = true
                toast_message = json.message
                bet_multiple = ""
                flag_btnbuy = true;
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
   
    const handleclick_infogame = () => {
        isModalInfo = true
    };
    const handleclick_listminbet = () => {
        isModalMinBet = true
    };
    const handle_minbet = (e) => {
        field_bet = e
        isModalMinBet = false
    };
    
    const call_reset = () => {
        bet_multiple = ""
        field_bet = engine_minbet
    };
    const call_allinvoice = () => {
        fetch_invoiceall()
    };
    const call_listresult = () => {
        fetch_listresult()
    };
   
    let nomor = [
        {id:"00",tipe:"RED",css:"btn-error"},
        {id:"01",tipe:"BLACK",css:""},
        {id:"02",tipe:"RED",css:"btn-error"},
        {id:"03",tipe:"BLACK",css:""},
        {id:"04",tipe:"RED",css:"btn-error"},
        {id:"05",tipe:"BLACK",css:""},
        {id:"06",tipe:"RED",css:"btn-error"},
        {id:"07",tipe:"BLACK",css:""},
        {id:"08",tipe:"RED",css:"btn-error"},
        {id:"09",tipe:"BLACK",css:""},
        {id:"10",tipe:"BLACK",css:""},
        {id:"11",tipe:"RED",css:"btn-error"},
        {id:"12",tipe:"BLACK",css:""},
        {id:"13",tipe:"RED",css:"btn-error"},
        {id:"14",tipe:"BLACK",css:""},
        {id:"15",tipe:"RED",css:"btn-error"},
        {id:"16",tipe:"BLACK",css:""},
        {id:"17",tipe:"RED",css:"btn-error"},
        {id:"18",tipe:"BLACK",css:""},
        {id:"19",tipe:"RED",css:"btn-error"},
        {id:"20",tipe:"RED",css:"btn-error"},
        {id:"21",tipe:"BLACK",css:""},
        {id:"22",tipe:"RED",css:"btn-error"},
        {id:"23",tipe:"BLACK",css:""},
        {id:"24",tipe:"RED",css:"btn-error"},
        {id:"25",tipe:"BLACK",css:""},
        {id:"26",tipe:"RED",css:"btn-error"},
        {id:"27",tipe:"BLACK",css:""},
        {id:"28",tipe:"RED",css:"btn-error"},
        {id:"29",tipe:"BLACK",css:""},
        {id:"30",tipe:"BLACK",css:""},
        {id:"31",tipe:"RED",css:"btn-error"},
        {id:"32",tipe:"BLACK",css:""},
        {id:"33",tipe:"RED",css:"btn-error"},
        {id:"34",tipe:"BLACK",css:""},
        {id:"35",tipe:"RED",css:"btn-error"},
        {id:"36",tipe:"BLACK",css:""},
        {id:"37",tipe:"RED",css:"btn-error"},
        {id:"38",tipe:"BLACK",css:""},
        {id:"39",tipe:"RED",css:"btn-error"},
        {id:"40",tipe:"RED",css:"btn-error"},
        {id:"41",tipe:"BLACK",css:""},
        {id:"42",tipe:"RED",css:"btn-error"},
        {id:"43",tipe:"BLACK",css:""},
        {id:"44",tipe:"RED",css:"btn-error"},
        {id:"45",tipe:"BLACK",css:""},
        {id:"46",tipe:"RED",css:"btn-error"},
        {id:"47",tipe:"BLACK",css:""},
        {id:"48",tipe:"RED",css:"btn-error"},
        {id:"49",tipe:"BLACK",css:""},
        {id:"50",tipe:"BLACK",css:""},
        {id:"51",tipe:"RED",css:"btn-error"},
        {id:"52",tipe:"BLACK",css:""},
        {id:"53",tipe:"RED",css:"btn-error"},
        {id:"54",tipe:"BLACK",css:""},
        {id:"55",tipe:"RED",css:"btn-error"},
        {id:"56",tipe:"BLACK",css:""},
        {id:"57",tipe:"RED",css:"btn-error"},
        {id:"58",tipe:"BLACK",css:""},
        {id:"59",tipe:"RED",css:"btn-error"},
        {id:"60",tipe:"RED",css:"btn-error"},
        {id:"61",tipe:"BLACK",css:""},
        {id:"62",tipe:"RED",css:"btn-error"},
        {id:"63",tipe:"BLACK",css:""},
        {id:"64",tipe:"RED",css:"btn-error"},
        {id:"65",tipe:"BLACK",css:""},
        {id:"66",tipe:"RED",css:"btn-error"},
        {id:"67",tipe:"BLACK",css:""},
        {id:"68",tipe:"RED",css:"btn-error"},
        {id:"69",tipe:"BLACK",css:""},
        {id:"70",tipe:"BLACK",css:""},
        {id:"71",tipe:"RED",css:"btn-error"},
        {id:"72",tipe:"BLACK",css:""},
        {id:"73",tipe:"RED",css:"btn-error"},
        {id:"74",tipe:"BLACK",css:""},
        {id:"75",tipe:"RED",css:"btn-error"},
        {id:"76",tipe:"BLACK",css:""},
        {id:"77",tipe:"RED",css:"btn-error"},
        {id:"78",tipe:"BLACK",css:""},
        {id:"79",tipe:"RED",css:"btn-error"},
        {id:"80",tipe:"RED",css:"btn-error"},
        {id:"81",tipe:"BLACK",css:""},
        {id:"82",tipe:"RED",css:"btn-error"},
        {id:"83",tipe:"BLACK",css:""},
        {id:"84",tipe:"RED",css:"btn-error"},
        {id:"85",tipe:"BLACK",css:""},
        {id:"86",tipe:"RED",css:"btn-error"},
        {id:"87",tipe:"BLACK",css:""},
        {id:"88",tipe:"RED",css:"btn-error"},
        {id:"89",tipe:"BLACK",css:""},
        {id:"90",tipe:"BLACK",css:""},
        {id:"91",tipe:"RED",css:"btn-error"},
        {id:"92",tipe:"BLACK",css:""},
        {id:"93",tipe:"RED",css:"btn-error"},
        {id:"94",tipe:"BLACK",css:""},
        {id:"95",tipe:"RED",css:"btn-error"},
        {id:"96",tipe:"BLACK",css:""},
        {id:"97",tipe:"RED",css:"btn-error"},
        {id:"98",tipe:"BLACK",css:""},
        {id:"99",tipe:"RED",css:"btn-error"},
    ]
  </script>
 
<section class="glass bg-opacity-60 rounded-lg">
    <section class="flex-col w-full p-2 rounded-md ">
        <center>
            <img class="w-[150px]" src="https://i.imgur.com/PNSe1ov.png" alt="" srcset="">
        </center>
        <section class="hidden lg:flex justify-between w-full bg-base-100 p-1 rounded-md select-none mt-1">
            <section class="flex-col text-center  font-bold  w-1/2  ">
                <div class="flex-col">
                    <div class="text-lg lg:text-xl">PERIODE</div>
                    <div class="link-accent text-sm lg:text-lg">{engine_invoice}</div>
                </div>
                <div class="flex-col mt-2">
                    <div class="text-lg lg:text-xl">WAKTU</div>
                    <div class="link-accent text-sm lg:text-lg">{engine_time} S </div>
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
                    <div class="link-accent text-lg">{engine_invoice}</div>
                </div>
                <div class="flex-col w-full text-center">
                    <div class="text-sm">WAKTU</div>
                    <div class="link-accent text-lg">{engine_time} S </div>
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
        <section class="grid grid-cols-1 w-full gap-2 mt-2">
            <div class="h-[350px] w-full overflow-auto">
                <div class="grid grid-cols-6 sm:grid-cols-10 md:grid-cols-10 xl:grid-cols-10 lg:grid-cols-10 gap-1 w-full">
                    {#each nomor as rec}
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="{rec.id}" />
                        <div class="swap-on btn btn-outline">{rec.id}</div>
                        <div class="swap-off btn  ">{rec.id}</div>
                    </label>
                    {/each}
                </div>
            </div>
            <div class="flex-col">
                <div on:click={() => {
                    handleclick_listminbet();
                  }} class="flex-col w-full h-[90px] justify-center bg-base-300 cursor-pointer rounded-lg">
                    <div class="w-full p-2   text-center ">
                        <div class="uppercase text-xs">Pilih Coin Bet :</div>
                        <div class="text-5xl link-accent">{new Intl.NumberFormat().format(field_bet)}</div>
                    </div>
                </div>
            </div>
            {#if engine_status == "OPEN"}
                {#if flag_btnbuy}
                    <div class="grid grid-cols-3 gap-2 w-full">
                        <button on:click={() => {
                                    handleclick_infogame();
                            }}  class="btn btn-info ">
                            Info 
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round" d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z" />
                            </svg>
                        </button>
                        <button on:click={() => {
                                    call_reset();
                            }}  class="btn btn-warning">
                            Reset 
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M16.023 9.348h4.992v-.001M2.985 19.644v-4.992m0 0h4.992m-4.993 0 3.181 3.183a8.25 8.25 0 0 0 13.803-3.7M4.031 9.865a8.25 8.25 0 0 1 13.803-3.7l3.181 3.182m0-4.991v4.99" />
                            </svg>
                        </button>
                        <button on:click={() => {
                                    call_bayar();
                            }}  class="btn btn-success">
                            Bayar 
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
                            </svg>
                        </button>
                    </div>
                {/if}
            {/if}
        </section>
    </section>
</section>
<section class="flex-col gap-2 mt-4 p-2 glass bg-opacity-60 rounded-md">
    <div class="flex justify-center gap-2">
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

<input type="checkbox" id="my-modal-information" class="modal-toggle" bind:checked={isModalMinBet}>
<div class="modal" on:click|self={()=>isModalMinBet = false}>
    <div class="modal-box relative w-11/12 max-w-lg h-1/2 lg:h-2/3 overflow-auto select-none">
        <label for="my-modal-information" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">COIN BET</h3>
        <div class="h-fit overflow-auto  mt-2" >
            <div class="grid grid-cols-3 lg:grid-cols-5 mt-5 gap-2 justify-self-center">
              {#each client_listbet as rec}
                <div on:click={() => {
                    handle_minbet(rec.id);
                  }} 
                  class="btn btn-xs lg:btn-sm btn-outline btn-success cursor-pointer">{new Intl.NumberFormat().format(rec.val)}</div>
              {/each}
            </div>
            
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-infogame" class="modal-toggle" bind:checked={isModalInfo}>
<div class="modal" on:click|self={()=>isModalMinBet = false}>
    <div class="modal-box relative w-11/12 max-w-lg h-1/2 lg:h-2/3 overflow-auto select-none">
        <label for="my-modal-infogame" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">INFO</h3>
        <div class="h-fit overflow-auto  mt-2" >
            <p class="text-xs">
                ASDASD
            </p>
        </div>
    </div>
</div>
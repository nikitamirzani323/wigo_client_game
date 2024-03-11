<script>
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";
  
    
    dayjs.extend(utc);
    dayjs.extend(timezone);
   
    export let path_api = "";
    export let engine_time = 0
    export let engine_invoice = ""
    export let engine_status = "LOCK"
    export let client_company = ""
    export let client_username = ""
    export let client_timezone = "Asia/Jakarta"
    export let client_name = "";
    export let client_ipaddress = "";
    export let client_listbet = [];
    export let client_credit = 0;
    export let engine_minbet = 0;
    export let engine_multiplier = 0;
    
    let flag_toast = false;
    let toast_message = "";
    
    let clockmachine = "";
    
    let flag_btnbuy = true;
    let field_maxlength_bet = length
    let field_bet = 0
    let field_nomor = ""
    
    let list_invoice = []
    let list_result = []
    let keranjang = [];
    let flag_listinvoice = true;
    let flag_listresult = false;
    let bet_multiple = []
    let isModalMinBet = false;
    let isModalInfo = false;
    
    let redblack = []
    let btn_red_css = "btn btn-error"
    let btn_red_flag = false
    let btn_black_css = "btn"
    let btn_black_flag = false
    let btn_ganjil_css = "btn"
    let btn_ganjil_flag = false
    let btn_genap_css = "btn"
    let btn_genap_flag = false

    function updateClock() {
      let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm:ss");
      clockmachine = endtime;
    }
  
    function toast_hidden() {
        flag_toast = false;
    }
    $: {
        setInterval(updateClock, 1000);
        field_bet = engine_minbet
        if(client_username != "" && client_company != ""){
            fetch_invoiceall()
        }
        
    }
    async function call_bayar() {
        keranjang = [];
        let flag = true;
        let msg_err = ""
        const mergeResult = [...bet_multiple, ...redblack];
        let total_bet_multiple = mergeResult.length
        let total_bayar = parseInt(total_bet_multiple)*parseInt(field_bet)
       
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
            console.log(mergeResult)
            
            for(let i=0;i<total_bet_multiple;i++){
                let tipebet = "ANGKA"
                if(mergeResult[i] == "RED" || mergeResult[i] == "BLACK"){
                    tipebet = "REDBLACK"
                }
                const data = {
                    tipebet: tipebet,
                    nomor: mergeResult[i],
                    bet: parseInt(field_bet),
                    multiplier: parseFloat(engine_multiplier)
                };
                keranjang = [data, ...keranjang];
            }

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
                fetch_invoiceall()
                flag_toast = true
                flag_btnbuy = true;
                toast_message = json.message
                call_reset()
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
                invoice_company: client_company.toLowerCase(),
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
                invoice_company: client_company.toLowerCase(),
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
        keranjang = [];
        bet_multiple = []
        redblack = []
        field_bet = engine_minbet

        btn_red_css = "btn btn-error"
        btn_red_flag = false
        btn_black_css = "btn"
        btn_black_flag = false
        btn_ganjil_css = "btn"
        btn_ganjil_flag = false
        btn_genap_css = "btn"
        btn_genap_flag = false
    };
    const call_allinvoice = () => {
        fetch_invoiceall()
    };
    const call_listresult = () => {
        fetch_listresult()
    };
    
    const handleclick_redblack = (e) => {
        switch(e){
            case "RED":
                if(btn_red_flag == false){
                    btn_red_css = "btn btn-outline"
                    btn_red_flag = true;
                    redblack.push(e)
                }else{
                    btn_red_css = "btn btn-error"
                    btn_red_flag = false
                    for (let i = 0; i < redblack.length; i++) { 
                        if (redblack[i] === "RED") { 
                            redblack.splice(i, 1); 
                        } 
                    }
                }
                break;
            case "BLACK":
                if(btn_black_flag == false){
                    btn_black_css = "btn btn-outline"
                    btn_black_flag = true;
                    redblack.push(e)
                }else{
                    btn_black_css = "btn"
                    btn_black_flag = false
                    for (let i = 0; i < redblack.length; i++) { 
                        if (redblack[i] === "BLACK") { 
                            redblack.splice(i, 1); 
                        } 
                    }
                }
                break;
            case "GANJIL":
                if(btn_ganjil_flag == false){
                    btn_ganjil_css = "btn btn-outline"
                    btn_ganjil_flag = true;
                    redblack.push(e)
                }else{
                    btn_ganjil_css = "btn"
                    btn_ganjil_flag = false
                    for (let i = 0; i < redblack.length; i++) { 
                        if (redblack[i] === "GANJIL") { 
                            redblack.splice(i, 1); 
                        } 
                    }
                }
                break;
            case "GENAP":
                if(btn_genap_flag == false){
                    btn_genap_css = "btn btn-outline"
                    btn_genap_flag = true;
                    redblack.push(e)
                }else{
                    btn_genap_css = "btn"
                    btn_genap_flag = false
                    for (let i = 0; i < redblack.length; i++) { 
                        if (redblack[i] === "GENAP") { 
                            redblack.splice(i, 1); 
                        } 
                    }
                }
                break;
        }
    };
    let nomor = [
        {id:"00",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"01",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"02",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"03",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"04",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"05",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"06",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"07",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"08",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"09",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"10",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"11",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"12",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"13",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"14",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"15",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"16",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"17",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"18",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"19",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"20",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"21",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"22",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"23",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"24",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"25",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"26",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"27",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"28",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"29",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"30",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"31",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"32",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"33",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"34",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"35",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"36",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"37",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"38",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"39",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"40",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"41",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"42",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"43",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"44",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"45",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"46",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"47",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"48",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"49",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"50",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"51",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"52",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"53",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"54",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"55",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"56",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"57",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"58",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"59",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"60",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"61",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"62",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"63",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"64",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"65",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"66",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"67",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"68",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"69",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"70",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"71",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"72",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"73",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"74",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"75",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"76",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"77",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"78",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"79",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"80",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"81",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"82",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"83",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"84",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"85",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"86",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"87",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"88",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"89",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"90",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"91",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"92",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"93",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"94",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"95",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"96",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"97",tipeangka:"GANJIL",tipe:"BLACK",css:""},
        {id:"98",tipeangka:"GENAP",tipe:"RED",css:"btn-error"},
        {id:"99",tipeangka:"GANJIL",tipe:"BLACK",css:""},
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
                <div class="grid grid-cols-4 gap-2">
                    <button  on:click={() => {
                        handleclick_redblack("GANJIL");
                    }} class="{btn_ganjil_css}">GANJIL</button>
                    <button  on:click={() => {
                        handleclick_redblack("BLACK");
                    }} class="{btn_black_css}">BLACK</button>
                    <button  on:click={() => {
                        handleclick_redblack("RED");
                    }} class="{btn_red_css}">RED</button>
                    <button  on:click={() => {
                        handleclick_redblack("GENAP");
                    }} class="{btn_genap_css}">GENAP</button>
                </div>
                <div class="grid grid-cols-6 mt-2 sm:grid-cols-10 md:grid-cols-10 xl:grid-cols-10 lg:grid-cols-10 gap-1 w-full">
                    {#each nomor as rec}
                    <label class="swap text-6xl">
                        <input bind:group={bet_multiple} type="checkbox" value="{rec.id}" />
                        <div class="swap-on btn btn-outline">{rec.id}</div>
                        <div class="swap-off btn {rec.css} ">{rec.id}</div>
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
                            <span class="{rec.invoiceclient_status_css} p-1 text-xs lg:text-sm  uppercase  rounded-lg w-20 text-black ">{rec.invoiceclient_status}</span>
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
                    handle_minbet(rec.money_bet);
                  }} 
                  class="btn btn-xs lg:btn-sm btn-outline btn-success cursor-pointer">{new Intl.NumberFormat().format(rec.money_bet)}</div>
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
            <table class="table table-xs">
                <tr>
                    <td width="30%">MIN BET</td>
                    <td width="1%">:</td>
                    <td width="*">{new Intl.NumberFormat().format(field_bet)}</td>
                </tr>
                <tr>
                    <td>HADIAH ANGKA (00 - 99)</td>
                    <td>:</td>
                    <td>x 5</td>
                </tr>
                <tr>
                    <td>HADIAH GENAP/GANJIL, RED/BLACK</td>
                    <td>:</td>
                    <td>x 0.95</td>
                </tr>
            </table>
            <p class="text-[12px] mt-2">
                <b class="uppercase font-bold">Cara Bermain :</b> <br />
                Pilih angka 00 - 99 <br />
                Nomor akan diundi setelah waktu 0 Second,  <br />
                jika nomor anda kena, maka anda akan mendapatkan: modal + (modal * 5)
                <br /><br />
                Contoh :<br />
                Anda memasang nomor 25, dengan bet 500<br />
                keluaran adalah nomor 25<br />
                jadi bet anda menang <br />
                pembayarannya adalah : modal + (modal x 5)<br />
                500 + (500 x 5) = 3000<br />
                anda akan mendapatkan 3000
                <br /><br />
                Anda memasang nomor 10, dengan bet 500<br />
                keluaran adalah nomor 00<br />
                jadi bet anda kalah <br />
        </div>
    </div>
</div>
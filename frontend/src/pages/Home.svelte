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
    let nomor_global = 0;
    let flag_listinvoice = true;
    let flag_listresult = false;
    let bet_multiple = ""
    let bet_multiple2 = ""
    let bet_multiple3 = ""
    let bet_multiple4 = ""

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
        let mergednomor = [...bet_multiple, ...bet_multiple2, ...bet_multiple3, ...bet_multiple4]
        let total_mergednomor = mergednomor.length
        let total_bayar = parseInt(total_mergednomor)*parseInt(field_bet)
        console.log(mergednomor.toString())
        console.log("total nomor : ",total_mergednomor)
        console.log("total bayar : ",total_bayar)

        if(parseInt(engine_time) < 5){
            flag = false
            msg_err = "Timeout"
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
        if(parseInt(total_bayar) > parseInt(client_credit)){
            flag = false
            msg_err = "Credit tidak cukup "
        }
        // flag = false;
        if(flag){
            flag_btnbuy = false;
            for(let i=0;i<total_mergednomor;i++){
                const data = {
                    nomor: mergednomor[i],
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
                bet_multiple2 = ""
                bet_multiple3 = ""
                bet_multiple4 = ""
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
   
    let flag_page1 = true;
    let flag_page2 = false;
    let flag_page3 = false;
    let flag_page4 = false;
   
    const call_pagenomor = (e) => {
        switch(e){
            case "page1":
                flag_page1 = true;
                flag_page2 = false;
                flag_page3 = false;
                flag_page4 = false;
                break;
            case "page2":
                flag_page1 = false;
                flag_page2 = true;
                flag_page3 = false;
                flag_page4 = false;
                break;
            case "page3":
                flag_page1 = false;
                flag_page2 = false;
                flag_page3 = true;
                flag_page4 = false;
                break;
            case "page4":
                flag_page1 = false;
                flag_page2 = false;
                flag_page3 = false;
                flag_page4 = true;
                break;
        }
    };
  
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
            <div class="hidden flex w-full bg-base-300">
                <input on:keyup={handleKeyboard_nomor}
                    bind:value={field_nomor}  
                    class="w-full text-[50px] text-center link-accent bg-base-300 focus:outline-none"
                    type="text" 
                    placeholder="2D" 
                    maxlength="2">
            </div>
            
            <div class="flex justify-center gap-2">
                <button on:click={() => {
                    call_pagenomor("page1");
                 }}  class="btn btn-xs btn-info"><img src="hdmi-fill.svg" ></button>
                <button on:click={() => {
                    call_pagenomor("page2");
                 }}  class="btn btn-xs btn-info"><img src="hdmi-fill.svg" ></button>
                <button on:click={() => {
                    call_pagenomor("page3");
                 }}  class="btn btn-xs btn-info"><img src="hdmi-fill.svg" ></button>
                 <button on:click={() => {
                    call_pagenomor("page4");
                 }}  class="btn btn-xs btn-info"><img src="hdmi-fill.svg" ></button>
            </div>
            {#if flag_page1}
                <div class="grid grid-cols-5 lg:grid-cols-5 gap-2">
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="00" />
                        <div class="swap-on btn btn-outline">00</div>
                        <div class="swap-off btn  ">00</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="01" />
                        <div class="swap-on btn btn-outline">01</div>
                        <div class="swap-off btn ">01</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="02" />
                        <div class="swap-on btn btn-outline">02</div>
                        <div class="swap-off btn ">02</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="03" />
                        <div class="swap-on btn btn-outline">03</div>
                        <div class="swap-off btn ">03</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="04" />
                        <div class="swap-on btn  btn-outline">04</div>
                        <div class="swap-off btn ">04</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="05" />
                        <div class="swap-on btn btn-outline">05</div>
                        <div class="swap-off btn ">05</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="06" />
                        <div class="swap-on btn btn-outline">06</div>
                        <div class="swap-off btn ">06</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="07" />
                        <div class="swap-on btn  btn-outline">07</div>
                        <div class="swap-off btn ">07</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="08" />
                        <div class="swap-on btn  btn-outline">08</div>
                        <div class="swap-off btn  ">08</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="09" />
                        <div class="swap-on btn btn-outline">09</div>
                        <div class="swap-off btn ">09</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="10" />
                        <div class="swap-on btn btn-outline">10</div>
                        <div class="swap-off btn ">10</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="11" />
                        <div class="swap-on btn btn-outline">11</div>
                        <div class="swap-off btn ">11</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="12" />
                        <div class="swap-on btn btn-outline">12</div>
                        <div class="swap-off btn ">12</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="13" />
                        <div class="swap-on btn btn-outline">13</div>
                        <div class="swap-off btn ">13</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="14" />
                        <div class="swap-on btn btn-outline">14</div>
                        <div class="swap-off btn ">14</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="15" />
                        <div class="swap-on btn btn-outline">15</div>
                        <div class="swap-off btn ">15</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="16" />
                        <div class="swap-on btn btn-outline">16</div>
                        <div class="swap-off btn ">16</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="17" />
                        <div class="swap-on btn btn-outline">17</div>
                        <div class="swap-off btn ">17</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="18" />
                        <div class="swap-on btn btn-outline">18</div>
                        <div class="swap-off btn ">18</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="19" />
                        <div class="swap-on btn btn-outline">19</div>
                        <div class="swap-off btn ">19</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="20" />
                        <div class="swap-on btn btn-outline">20</div>
                        <div class="swap-off btn ">20</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="21" />
                        <div class="swap-on btn btn-outline">21</div>
                        <div class="swap-off btn ">21</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="22" />
                        <div class="swap-on btn btn-outline">22</div>
                        <div class="swap-off btn ">22</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="23" />
                        <div class="swap-on btn btn-outline">23</div>
                        <div class="swap-off btn ">23</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple} type="checkbox" value="24" />
                        <div class="swap-on btn btn-outline">24</div>
                        <div class="swap-off btn ">24</div>
                    </label>
                </div>
            {/if}
            {#if flag_page2}
                <div class="grid grid-cols-5 lg:grid-cols-5 gap-2">
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="25" />
                        <div class="swap-on btn btn-outline">25</div>
                        <div class="swap-off btn ">25</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="26" />
                        <div class="swap-on btn btn-outline">26</div>
                        <div class="swap-off btn ">26</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="27" />
                        <div class="swap-on btn btn-outline">27</div>
                        <div class="swap-off btn ">27</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="28" />
                        <div class="swap-on btn btn-outline">28</div>
                        <div class="swap-off btn ">28</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="29" />
                        <div class="swap-on btn btn-outline">29</div>
                        <div class="swap-off btn ">29</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="30" />
                        <div class="swap-on btn btn-outline">30</div>
                        <div class="swap-off btn ">30</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="31" />
                        <div class="swap-on btn btn-outline">31</div>
                        <div class="swap-off btn ">31</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="32" />
                        <div class="swap-on btn btn-outline">32</div>
                        <div class="swap-off btn ">32</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="33" />
                        <div class="swap-on btn btn-outline">33</div>
                        <div class="swap-off btn ">33</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="34" />
                        <div class="swap-on btn btn-outline">34</div>
                        <div class="swap-off btn ">34</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="35" />
                        <div class="swap-on btn btn-outline">35</div>
                        <div class="swap-off btn ">35</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="36" />
                        <div class="swap-on btn btn-outline">36</div>
                        <div class="swap-off btn ">36</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="37" />
                        <div class="swap-on btn btn-outline">37</div>
                        <div class="swap-off btn ">37</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="38" />
                        <div class="swap-on btn btn-outline">38</div>
                        <div class="swap-off btn ">38</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="39" />
                        <div class="swap-on btn btn-outline">39</div>
                        <div class="swap-off btn ">39</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="40" />
                        <div class="swap-on btn btn-outline">40</div>
                        <div class="swap-off btn ">40</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="41" />
                        <div class="swap-on btn btn-outline">41</div>
                        <div class="swap-off btn ">41</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="42" />
                        <div class="swap-on btn btn-outline">42</div>
                        <div class="swap-off btn ">42</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="43" />
                        <div class="swap-on btn btn-outline">43</div>
                        <div class="swap-off btn ">43</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="44" />
                        <div class="swap-on btn btn-outline">44</div>
                        <div class="swap-off btn ">44</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="45" />
                        <div class="swap-on btn btn-outline">45</div>
                        <div class="swap-off btn ">45</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="46" />
                        <div class="swap-on btn btn-outline">46</div>
                        <div class="swap-off btn ">46</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="47" />
                        <div class="swap-on btn btn-outline">47</div>
                        <div class="swap-off btn ">47</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="48" />
                        <div class="swap-on btn btn-outline">48</div>
                        <div class="swap-off btn ">48</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple2} type="checkbox" value="49" />
                        <div class="swap-on btn btn-outline">49</div>
                        <div class="swap-off btn ">49</div>
                    </label>
                </div>
            {/if}
            {#if flag_page3}
                <div class="grid grid-cols-5 lg:grid-cols-5 gap-2">
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="50" />
                        <div class="swap-on btn btn-outline">50</div>
                        <div class="swap-off btn ">50</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="51" />
                        <div class="swap-on btn btn-outline">51</div>
                        <div class="swap-off btn ">51</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="52" />
                        <div class="swap-on btn btn-outline">52</div>
                        <div class="swap-off btn ">52</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="53" />
                        <div class="swap-on btn btn-outline">53</div>
                        <div class="swap-off btn ">53</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="54" />
                        <div class="swap-on btn btn-outline">54</div>
                        <div class="swap-off btn ">54</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="55" />
                        <div class="swap-on btn btn-outline">55</div>
                        <div class="swap-off btn ">55</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="56" />
                        <div class="swap-on btn btn-outline">56</div>
                        <div class="swap-off btn ">56</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="57" />
                        <div class="swap-on btn btn-outline">57</div>
                        <div class="swap-off btn ">57</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="58" />
                        <div class="swap-on btn btn-outline">58</div>
                        <div class="swap-off btn ">58</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="59" />
                        <div class="swap-on btn btn-outline">59</div>
                        <div class="swap-off btn ">59</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="60" />
                        <div class="swap-on btn btn-outline">60</div>
                        <div class="swap-off btn ">60</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="61" />
                        <div class="swap-on btn btn-outline">61</div>
                        <div class="swap-off btn ">61</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="62" />
                        <div class="swap-on btn btn-outline">62</div>
                        <div class="swap-off btn ">62</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="63" />
                        <div class="swap-on btn btn-outline">63</div>
                        <div class="swap-off btn ">63</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="64" />
                        <div class="swap-on btn btn-outline">64</div>
                        <div class="swap-off btn ">64</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="65" />
                        <div class="swap-on btn btn-outline">65</div>
                        <div class="swap-off btn ">65</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="66" />
                        <div class="swap-on btn btn-outline">66</div>
                        <div class="swap-off btn ">66</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="67" />
                        <div class="swap-on btn btn-outline">67</div>
                        <div class="swap-off btn ">67</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="68" />
                        <div class="swap-on btn btn-outline">68</div>
                        <div class="swap-off btn ">68</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="69" />
                        <div class="swap-on btn btn-outline">69</div>
                        <div class="swap-off btn ">69</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="70" />
                        <div class="swap-on btn btn-outline">70</div>
                        <div class="swap-off btn ">70</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="71" />
                        <div class="swap-on btn btn-outline">71</div>
                        <div class="swap-off btn ">71</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="72" />
                        <div class="swap-on btn btn-outline">72</div>
                        <div class="swap-off btn ">72</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="73" />
                        <div class="swap-on btn btn-outline">73</div>
                        <div class="swap-off btn ">73</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple3} type="checkbox" value="74" />
                        <div class="swap-on btn btn-outline">74</div>
                        <div class="swap-off btn ">74</div>
                    </label>
                </div>
            {/if}
            {#if flag_page4}
                <div class="grid grid-cols-5 lg:grid-cols-5 gap-2">
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="75" />
                        <div class="swap-on btn btn-outline">75</div>
                        <div class="swap-off btn ">75</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="76" />
                        <div class="swap-on btn btn-outline">76</div>
                        <div class="swap-off btn ">76</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="77" />
                        <div class="swap-on btn btn-outline">77</div>
                        <div class="swap-off btn ">77</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="78" />
                        <div class="swap-on btn btn-outline">78</div>
                        <div class="swap-off btn ">78</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="79" />
                        <div class="swap-on btn btn-outline">79</div>
                        <div class="swap-off btn ">79</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="80" />
                        <div class="swap-on btn btn-outline">80</div>
                        <div class="swap-off btn ">80</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="81" />
                        <div class="swap-on btn btn-outline">81</div>
                        <div class="swap-off btn ">81</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="82" />
                        <div class="swap-on btn btn-outline">82</div>
                        <div class="swap-off btn ">82</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="83" />
                        <div class="swap-on btn btn-outline">83</div>
                        <div class="swap-off btn ">83</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="84" />
                        <div class="swap-on btn btn-outline">84</div>
                        <div class="swap-off btn ">84</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="85" />
                        <div class="swap-on btn btn-outline">60</div>
                        <div class="swap-off btn ">60</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="86" />
                        <div class="swap-on btn btn-outline">61</div>
                        <div class="swap-off btn ">61</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="87" />
                        <div class="swap-on btn btn-outline">62</div>
                        <div class="swap-off btn ">62</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="88" />
                        <div class="swap-on btn btn-outline">88</div>
                        <div class="swap-off btn ">88</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="89" />
                        <div class="swap-on btn btn-outline">89</div>
                        <div class="swap-off btn ">89</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="90" />
                        <div class="swap-on btn btn-outline">90</div>
                        <div class="swap-off btn ">90</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="91" />
                        <div class="swap-on btn btn-outline">91</div>
                        <div class="swap-off btn ">91</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="92" />
                        <div class="swap-on btn btn-outline">92</div>
                        <div class="swap-off btn ">92</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="93" />
                        <div class="swap-on btn btn-outline">93</div>
                        <div class="swap-off btn ">93</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="94" />
                        <div class="swap-on btn btn-outline">94</div>
                        <div class="swap-off btn ">94</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="95" />
                        <div class="swap-on btn btn-outline">95</div>
                        <div class="swap-off btn ">95</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="96" />
                        <div class="swap-on btn btn-outline">96</div>
                        <div class="swap-off btn ">96</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="97" />
                        <div class="swap-on btn btn-outline">97</div>
                        <div class="swap-off btn ">97</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="98" />
                        <div class="swap-on btn btn-outline">98</div>
                        <div class="swap-off btn ">98</div>
                    </label>
                    <label class="swap">
                        <input bind:group={bet_multiple4} type="checkbox" value="99" />
                        <div class="swap-on btn btn-outline">99</div>
                        <div class="swap-off btn ">99</div>
                    </label>
                </div>
            {/if}
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
                {#if flag_btnbuy}
                    <button on:click={() => {
                                call_bayar();
                        }}  class="btn btn-success">
                        Bayar 
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 0 0-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 0 0-16.536-1.84M7.5 14.25 5.106 5.272M6 20.25a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Zm12.75 0a.75.75 0 1 1-1.5 0 .75.75 0 0 1 1.5 0Z" />
                        </svg>
                    </button>
                {/if}
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
  
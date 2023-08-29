const postamatInputs = document.getElementById("postamatInputs");
const courierInputs = document.getElementById("courierInputs");
const departmentInputs = document.getElementById("departmentInputs");

const deliveryMethodRadio = document.querySelectorAll('input[name="delivery"]');

const selectCityDepartment = document.getElementById("selectCityDepartment")
const selectCityWarehouse = document.getElementById("selectCityWarehouse")

const selectWarehouse = document.getElementById("selectWarehouse")
const selectWarehouseSettlement = document.getElementById("selectWarehouseSettlement")

const schedule = document.getElementById("schedule")

const arePostamat = document.getElementById("arePostamat")

deliveryMethodRadio.forEach(radio => {
    radio.addEventListener("change", () => {
        if (radio.value === "postamat") {
            postamatInputs.style.display = "block";
            courierInputs.style.display = "none";
            departmentInputs.style.display = "none";
        } else if (radio.value === "courier") {
            postamatInputs.style.display = "none";
            courierInputs.style.display = "block";
            departmentInputs.style.display = "none";
        } else if (radio.value === "department") {
            postamatInputs.style.display = "none";
            courierInputs.style.display = "none";
            departmentInputs.style.display = "block";
        }
    });
});

function FindCity(e,deliveryMethod) {
    const searchWord = e.value;

    const formData = new FormData();
    formData.append("searchWord", searchWord);

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/order/findCityNovaPoshta");
    xhr.send(formData);

    xhr.onload = function() {
        if (xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            if (deliveryMethod === "department"){
                DepartmentFindCity(response)
            }else if (deliveryMethod === "postamat"){
                arePostamat.style.display = "none"
                WarehouseFindCity(response)
            }
        } else {
            console.log(xhr.response)
        }
    };
}

function DepartmentFindCity(response){
    selectCityDepartment.innerHTML = ""

    for (const data of response.data) {
        selectCityDepartment.append(OptionCity(data))
    }

    FindSettlement(selectCityDepartment)
}

function WarehouseFindCity(response){
    selectCityWarehouse.innerHTML = ""

    for (const data of response.data) {
        selectCityWarehouse.append(OptionCity(data))
    }

    FindWarehouses(selectCityWarehouse)
}

function OptionCity(data) {
    console.log(data)
    let option = document.createElement("option")
    option.value = data.Ref
    option.innerText = data.SettlementTypeDescription + " " + data.Description + " " + data.AreaDescription + " обл"

    return option
}

function FindWarehouses(e){
    const formData = new FormData();
    formData.append("cityRef", e.value);

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/order/getWarehouses");
    xhr.send(formData);

    xhr.onload = function() {
        if (xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log(response.data)
            Warehouses(response)
        } else {
            console.log(xhr.response)
        }
    };
}

function Warehouses(response){
    selectWarehouse.innerHTML = ""
    selectWarehouse.data = response.data

    let arePostomatInCity = false

    for (const data of response.data) {
        if (data.CategoryOfWarehouse === "Postomat"){
            arePostomatInCity = true
            selectWarehouse.append(OptionWarehouse(data))
        }
    }

    if (!arePostomatInCity){
        arePostamat.style.display = "block"
    }
}

function OptionWarehouse(data) {
    let option = document.createElement("option")
    option.value = data.Ref
    option.innerText = data.Description

    return option
}

function FindSettlement(e){
    const formData = new FormData();
    formData.append("cityRef", e.value);

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/order/GetSettlements");
    xhr.send(formData);

    xhr.onload = function() {
        if (xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            Settlement(response)

            Schedule(selectWarehouse)
        } else {
            console.log(xhr.response)
        }
    };
}

function Settlement(response){
    selectWarehouseSettlement.innerHTML = ""
    selectWarehouseSettlement.data = response.data

    for (const data of response.data) {
        // if (data.CategoryOfWarehouse === "Branch"){
            selectWarehouseSettlement.append(OptionSettlement(data))
        // }
    }
}

function OptionSettlement(data) {
    console.log(data)
    let option = document.createElement("option")
    option.value = data.Ref
    option.innerText = data.Description

    return option
}

function Schedule(e){
    const warehouse = e.data.find(warehouse => warehouse.Ref === e.value);

    schedule.innerHTML = ""

    for (const day in warehouse.Schedule) {
        p = document.createElement("p")
        p.innerText = day + " " + warehouse.Schedule[day]
        schedule.append(p)
    }
}
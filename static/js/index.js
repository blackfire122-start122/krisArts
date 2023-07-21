function findInput(e){
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {

                let response = xhr.responseText;
                console.log(response);


            } else {
                console.error('Помилка ' + xhr.status + ': ' + xhr.statusText);
            }
        }
    };

    xhr.open('GET', 'localhost:8080/api/findArts', true);
    xhr.send();
}
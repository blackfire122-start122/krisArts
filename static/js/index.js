const ulElement = document.querySelector('.listArts');

const findInput = document.querySelector('.findInput');

function findArts() {
  let xhr = new XMLHttpRequest();

  xhr.onreadystatechange = function() {
    if (xhr.readyState === 4) {
      if (xhr.status === 200) {
        let response = JSON.parse(xhr.responseText);

        ulElement.innerHTML = '';
        
        response.forEach(artData => {
          const liElement = createArtElement(artData);
          ulElement.appendChild(liElement);
        });

      } else {
        console.error('Помилка ' + xhr.status + ': ' + xhr.statusText);
      }
    }
  };
  const url = `http://localhost:8080/api/findArts?find=${encodeURIComponent(findInput.value)}`;

  xhr.open('GET', url, true);
  xhr.send();
}

function createArtElement(artData) {
  const liElement = document.createElement('li');
  
  liElement.innerHTML = `
    <h3>${artData.Name}</h3>
    <img src="${artData.Image}" alt="Artwork" width="200" height="200">
    <p>Description: ${artData.Description}</p>
    <p>Price: ${artData.Price}</p>
  `;

  return liElement;
}

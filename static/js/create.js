const fileInput = document.getElementById('image');
const fileBtn = document.getElementById('file-btn');
const fileName = document.getElementById('file-name');

fileInput.addEventListener('change',
  function() {
    const file = fileInput.files[0];
    if (file) {
      fileName.textContent = file.name;
    } else {
      fileName.textContent = 'Файл не вибрано';
    }
  });

  fileBtn.addEventListener('click', function() {
     fileInput.click();
  });
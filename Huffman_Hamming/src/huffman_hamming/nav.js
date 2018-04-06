var cadena = "";
var indice = 0;
if (id== 0){
  cadena += `<li class="nav-item">
                <a class="nav-link disabled">No existen diferencias</a>
              </li>`;
}
for (;indice<id;indice++){
  cadena += `<li class="nav-item">
                <a class="nav-link" href="#id`+indice+`">`+(indice+1)+`</a>
              </li>`;
  }
document.write(cadena);

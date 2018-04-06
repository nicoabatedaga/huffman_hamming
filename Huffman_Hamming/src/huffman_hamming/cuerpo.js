dmp.diff_cleanupSemantic(diff);
var resultado="";
var i =0;
var id=0;
for (;i<diff.length;i++){
    switch (diff[i][0]) {
      case -1:
            resultado+='<p class="focus text-danger" id="id'+id+'" style="background:#ffe6e6;">-'+diff[i][1]+'</p>';
            id++;
        break;
        case 0:
            resultado+=diff[i][1];

          break;
          case 1:
            resultado+='<p class="focus text-success" id="id'+id+'"   style="background:#e6ffe6;">+'+diff[i][1]+'</p>';
            id++;
            break;
      default:
      resultado+='\n';
    }
  }
  resultado = resultado.split("\n").join("<br>");
  document.getElementById('cuerpo').innerHTML =resultado;

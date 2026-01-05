function parseNums(str){
  return str.split(',').map(s=>parseFloat(s.trim())).filter(n=>!isNaN(n));
}

function average(arr){
  if(!arr.length) return 0;
  return arr.reduce((a,b)=>a+b,0)/arr.length;
}

function median(arr){
  if(!arr.length) return 0;
  const s = [...arr].sort((a,b)=>a-b);
  const m = Math.floor(s.length/2);
  return (s.length%2) ? s[m] : (s[m-1]+s[m])/2;
}

function variance(arr){
  const mu = average(arr);
  return average(arr.map(x=> (x-mu)**2 ));
}

function sdev(arr){
  return Math.sqrt(variance(arr));
}

function compute(){
  const vals = parseNums(document.getElementById('nums').value);
  const out = document.getElementById('out');
  
  if(vals.length === 0) {
    out.innerHTML = '<div class="result-item"><span class="result-label">Please enter valid numbers</span></div>';
    return;
  }
  
  out.innerHTML = `
    <div class="result-item">
      <span class="result-label">Count:</span>
      <span class="result-value">${vals.length}</span>
    </div>
    <div class="result-item">
      <span class="result-label">Average:</span>
      <span class="result-value">${average(vals).toFixed(6)}</span>
    </div>
    <div class="result-item">
      <span class="result-label">Median:</span>
      <span class="result-value">${median(vals).toFixed(6)}</span>
    </div>
    <div class="result-item">
      <span class="result-label">Variance:</span>
      <span class="result-value">${variance(vals).toFixed(6)}</span>
    </div>
    <div class="result-item">
      <span class="result-label">Std Dev:</span>
      <span class="result-value">${sdev(vals).toFixed(6)}</span>
    </div>
  `;
}
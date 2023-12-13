const url =
  "https://glacier-api.avax.network/v1/networks/mainnet/validators?pageSize=100&validationStatus=active&subnetId=11111111111111111111111111111111LpoYY";

const validators = [];

let results = await fetch(url);
let data = await results.json();
validators.push(data.validators);
// console.log(data);
let nextPage = data.nextPageToken;
while (nextPage != null) {
  results = await fetch(`${url}&pageToken=${nextPage}`);
  data = await results.json();
  validators.push(data.validators);
  nextPage = data.nextPageToken;
  // console.log(nextPage);
}
console.log(JSON.stringify(validators.flat()));

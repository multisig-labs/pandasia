import { StandardMerkleTree } from "https://esm.sh/@openzeppelin/merkle-tree";

// const values = [
// 	["P-avax19zfygxaf59stehzedhxjesads0p5jdvfeedal0"], // 0x2892441ba9a160bcdc596dcd2cc3ad83c3493589
// 	["P-avax1adfqcxchsp3nnjj3a0jj3psgzs63ldggzew7c9"], // 0xeb520c1b17806339ca51ebe528860814351fb508
// 	["P-avax1gfpj30csekhwmf4mqkncelus5zl2ztqzvv7aww"], // 0x424328bf10cdaeeda6bb05a78cff90a0bea12c02
// 	["P-avax1tnuesf6cqwnjw7fxjyk7lhch0vhf0v95wj5jvy"], // 0x5cf998275803a7277926912defdf177b2e97b0b4
// ];

const values = [
	["0x1111111111111111111111111111111111111111"],
	["0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"],
];

const tree = StandardMerkleTree.of(values, ["address"]);
console.log("Merkle Root:", tree.root);
console.log(JSON.stringify(tree.dump(), null, 2));
console.log(JSON.stringify(tree.dump()));

const proof = tree.getProof(["0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"]);
console.log("Proof:", proof);
const verified = StandardMerkleTree.verify(
	tree.root,
	["address"],
	["0x0961Ca10D49B9B8e371aA0Bcf77fE5730b18f2E4"],
	proof
);

console.log("Verified:", verified);

// Merkle Root: 0x1261849b132545d29a7685fd7046d6200577cd912dda7f56b5a0d6dc16cb220d
// {
//   "format": "standard-v1",
//   "tree": [
//     "0x1261849b132545d29a7685fd7046d6200577cd912dda7f56b5a0d6dc16cb220d",
//     "0x76b6437b39eb2e64dca0afc556654206c04fe45f2f779295a04526d4facaf34f",
//     "0x3548906981a17dc8fedb3df19439a12bd993dbd555e0fb9e9a6eb967ee4401d0",
//     "0xa52e4f15acf7ff84255187ffb785366ccc84195823bcbd8760bf1b857cfe0b28",
//     "0x62ae3d5468d7a5dc6fcbc9597e8d38755e36ca08d189e51a4d9dfaab4c39746a",
//     "0x4ebd6c0d9a8b9bb3495d47825b1171aeb966d8593a0a995994f2b4e0167bba8c",
//     "0x20b2f891eaf390d96349554ee7297a8a8972c13215d1aa9dd752f9c6822c1888"
//   ],
//   "values": [
//     {
//       "value": [
//         "0x2892441ba9a160bcdc596dcd2cc3ad83c3493589"
//       ],
//       "treeIndex": 3
//     },
//     {
//       "value": [
//         "0xeb520c1b17806339ca51ebe528860814351fb508"
//       ],
//       "treeIndex": 4
//     },
//     {
//       "value": [
//         "0x424328bf10cdaeeda6bb05a78cff90a0bea12c02"
//       ],
//       "treeIndex": 6
//     },
//     {
//       "value": [
//         "0x5cf998275803a7277926912defdf177b2e97b0b4"
//       ],
//       "treeIndex": 5
//     }
//   ],
//   "leafEncoding": [
//     "address"
//   ]
// }

//this makes possible to diff between capitalized and normal letters


// export function duplicateEncode(word: string){
//   //const newWord = [...word].map((ch) => (word.indexOf(ch) != -1) ? "(" : ")")
//     const newWord = [...word].map((ch) => {
//       const reg = new RegExp(ch, "g")
//       const matches = word.match(reg)
//       if (matches != null)  {
//         return (matches.length > 1) ? ")" : "("
//       }
//     })
//   return newWord.join("")
// }


// export function duplicateEncodeWithoutCapitalLetters(word: string){
//     //const newWord = [...word].map((ch) => (word.indexOf(ch) != -1) ? "(" : ")")
//       const newWord = [...word].map((ch) => {
//         const reg = new RegExp("[" + ch.toLowerCase() + ch.toUpperCase() + "]")
//         // const reg = new RegExp(ch.toLowerCase() + "|" + ch.toUpperCase(), "g")
//         const matches = word.match(reg)
//         if (matches != null)  {
//           return (matches.length > 1) ? ")" : "("
//         }
//       })
//     return newWord.join("")
// }

export function duplicateEncode(word: string){
  return [...word].map((ch) => {

    const matches = word.match(new RegExp("[" + ch.toLowerCase() + ch.toUpperCase() + "]", "g"))
    if (matches != null)  {
      return (matches.length > 1) ? ")" : "("
    }
  }).join("")

}
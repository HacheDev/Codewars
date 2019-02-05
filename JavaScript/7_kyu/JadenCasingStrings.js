String.prototype.toJadenCase = function () {

    return this.split(' ').map(function(word)  {
      return word.replace(word.charAt(0), word.charAt(0).toUpperCase());
    }).join(' ');
    
  };
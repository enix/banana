
function extendArray() {
  Array.prototype.upsert = function(newElem, comparisonFn) {
    let found = false;

    const out = this.map(elem => {
      if (comparisonFn(elem)) {
        found = true;
        return newElem;
      }
  
      return elem;
    });

    if (!found) {
      out.push(newElem);
    }

    return out;
  }
};

export default extendArray;

$global["$github.com/molizz/psd$"] = {
  decodePackBitsLines: function (buf, lines, large) {
    var r = new Uint32Array(lines);
    var lens = new DataView(buf.buffer);
    if (large) {
      for (var i = 0; i < lines; ++i) {
        r[i] = lens.getUint32(i << 2, false);
      }
    } else {
      for (var i = 0; i < lines; ++i) {
        r[i] = lens.getUint16(i << 1, false);
      }
    }
    return r;
  },
  decodePackBits: function (dest, buf, lens) {
    var ofs = 0,
      d = 0;
    for (var i = 0; i < lens.length; ++i) {
      for (var j = 0, ln = lens[i]; j < ln; ) {
        if (buf[ofs] <= 0x7f) {
          var l = buf[ofs++] + 1;
          for (var k = 0; k < l; ++k) {
            dest[d++] = buf[ofs++];
          }
          j += l + 1;
          continue;
        }
        if (buf[ofs] == 0x80) {
          ofs++;
          continue;
        }
        for (var k = 0, l = 256 - buf[ofs++] + 1, c = buf[ofs++]; k < l; ++k) {
          dest[d++] = c;
        }
        j += 2;
      }
    }
    return true;
  },
};

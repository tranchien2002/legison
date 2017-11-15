module REGEX
  SIGN_FILTER = /ký\)(\s)+(((\p{Lu}\p{Ll}+)\s*)+(\p{Lu}\p{Ll}+)?)/
  BASEON_FILTER = /\nCăn cứ (.+)/
  CHAPTER_FILTER = /\n(CHƯƠNG|Chương)\s(M{1,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|C?D|D?C{1,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|X?L|L?X{1,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|I?V|V?I{1,3}))\s+(.+)/
  PARTICLE_FILTER = /\n.iê?.(u|ù)\s+(\d+)\.\s+(.+)/

end

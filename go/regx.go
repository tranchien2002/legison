package main

const ID_ENFORCER_REGX = `(S|s)ố:\s*(\d+\/\d+\/([A-Z][A-Z])(\d+))`

const BASE_ON_REGX = `\nCăn cứ (.+)`

const CHAPTER_REGX = `\n(Chương|CHƯƠNG)\s(M{1,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|C?D|D?C{1,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|X?L|L?X{1,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|I?V|V?I{1,3}))\s+(.+)|(\.\/\.)`

const TOPIC_REGX = `\n(Mục|MỤC)\s*(\d+)\s*(.+)`

const ARTICLE_REGX = `(\n.iê?.(u|ù)\s+(\d+)\.\s+(.+))|(\.\/\.)`

const EFFECTIVE_DATE_REGX = `này.*thi\s*hành.*ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4})`

const PASS_DATE_REGX = `thông\s*qua\s*ngày\s*(\d{1,2})\s*tháng\s*(\d{1,2})\s*năm\s*(\d{4})`

const SIGN_REGX = `ký\)(\s)+(((\p{Lu}\p{Ll}+)\s*)+(\p{Lu}\p{Ll}+)?)`
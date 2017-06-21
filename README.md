gonzh --- 阿拉伯数字转中文，支持金额和大小写 v1.0
========================

Inspired by [nzh](https://github.com/cnwhy/nzh)

## Features
1. 阿拉伯数字转中文大小写；
2. 阿拉伯数字转中文金额大小写，精确到分；
3. 支持负数和小数的转换；
4. 支持科学计数法和超大数字转换。

## 1. 阿拉伯数字转中文，支持大小写两种类型

```
// 五千四百三十二万一千九百五十八
fmt.Println(gonzh.EncodeString(EncodeString("54321958", false, false))

// 伍仟肆佰叁拾贰万壹仟玖佰伍拾捌
fmt.Println(gonzh.EncodeString("54321958", false, true))
```

## 2. 阿拉伯数字转中文金额，精确到分

```
// 叁圆壹角肆分
fmt.Println(gonzh.EncodeString("3.1415926", true, true))

// 壹仟叁佰伍拾陆圆整
fmt.Println(gonzh.EncodeString("1356", true, true))
```

## 3. 负数和小数转换

```
// 负八千九百七十六
fmt.Println(gonzh.EncodeString("-8976", false, false))

// 负捌仟玖佰柒拾陆圆整
fmt.Println(gonzh.EncodeString("-8976", true, true))

// 九千零八十万七千零六十点一零二零三"
fmt.Println(gonzh.EncodeString("090807060.102030", false, false))

// 一十二载三千四百五十六正七千八百九十八涧七千六百五十四沟三千二百一十二穰三千四百五十六秭七千八百九十八垓七千六百五十四京三千二百一十二兆三千四百五十六亿七千八百九十八万七千六百五十四
fmt.Println(gonzh.EncodeString("1234567898765432123456789876543212345678987654", false, false))
```

## 4. 科学计数法和超大数字转换

```
// 一十二点三四五
fmt.Println(gonzh.EncodeString("1234.5e-2", false, false))

// 零点零零零零零零零零零三一四一五九二六
fmt.Println(gonzh.EncodeString("3.1415926e-10", false, false))

// 一十二垓三千四百五十六京七千八百九十兆
fmt.Println(gonzh.EncodeString("1.23456789e+21", false, false))
```

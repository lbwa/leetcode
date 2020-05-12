# 20 easy 有效括号

给定一个只包括 `'('`，`')'`，`'{'`，`'}'`，`'['`，`']'`  的字符串，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。

1. 左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

## Tips

1. An interesting property about a valid parenthesis expression is that a sub-expression of a valid expression should also be a valid expression. (Not every sub-expression) e.g.

   ```
   { { } [ ] [ [ [ ] ] ] } is VALID expression
             [ [ [ ] ] ]    is VALID sub-expression
     { } [ ]                is VALID sub-expression
   ```

   Can we exploit this recursive structure somehow?

1. What if whenever we encounter a matching pair of parenthesis in the expression, we simply remove it from the expression? This would keep on shortening the expression. e.g.

   ```
   { { ( { } ) } }
         |_|

   { { (      ) } }
       |______|

   { {          } }
     |__________|

   {                }
   |________________|
    VALID EXPRESSION!
   ```

1. The **stack** data structure can come in handy here in representing this recursive structure of the problem. We can't really process this from the inside out because we don't have an idea about the overall structure. But, the stack can help us process this recursively i.e. from outside to inwards.

## source

- [leatcode-cn](https://leetcode-cn.com/problems/valid-parentheses/)

export function lessThan<E>(a: E, b: E) {
  return a < b
}

export function greaterThan<E>(a: E, b: E) {
  return a > b
}

export function swap<E, I extends number>(list: E[], a: I, b: I) {
  ;[list[a], list[b]] = [list[b], list[a]]
}

export const debounce = (timeout: number, cb: Function) => {
    let fn: number;
    return (...params: unknown[]) => {
        if (fn !== null) {
            clearTimeout(fn)
        }
        fn = setTimeout(() => {
            cb(...params);
        }, timeout)
    }
}
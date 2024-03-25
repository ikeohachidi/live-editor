export const throttle = (delay: number, cb: Function) => {
    let prevTime = new Date().getTime();

    return (...params: unknown[]) => {
        let now = new Date();

        if ((prevTime + delay) < now.getTime()) {
            cb(...params)
            prevTime = now.getTime();
        }
    }
}

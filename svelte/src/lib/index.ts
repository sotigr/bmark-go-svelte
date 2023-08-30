// place files you want to import through the `$lib` alias in this folder.

export function humanFileSize(bytes: number, si = false, dp = 1) {
    const thresh = si ? 1000 : 1024;

    if (Math.abs(bytes) < thresh) {
        return bytes + ' B';
    }

    const units = si
        ? ['kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
        : ['KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];
    let u = -1;
    const r = 10 ** dp;

    do {
        bytes /= thresh;
        ++u;
    } while (Math.round(Math.abs(bytes) * r) / r >= thresh && u < units.length - 1);


    return bytes.toFixed(dp) + ' ' + units[u];
}


export function humanDate(dt: Date): string {
    let d = new Date(dt)
    let mm = d.getMonth() + 1; // getMonth() is zero-based
    let dd = d.getDate();

    let day = (dd > 9 ? '' : '0') + dd
    let month = (mm > 9 ? '' : '0') + mm
    let year = d.getFullYear().toString().substring(2)

    return `${day}/${month}/${year}`

}
 
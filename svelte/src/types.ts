export type File = {
    name: string
    fullName: string
    size: number
    created: Date
    contentType: string
}

export type Dir = {
    name: string
    fullName: string
}

export type DirectoryListing = {
    files: File[]
    folders: Dir[]
}

declare global {
    interface Window { lazyLoad: any; }
}

window.lazyLoad = window.lazyLoad || {};
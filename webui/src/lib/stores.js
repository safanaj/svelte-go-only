import { writable } from 'svelte/store';
export const controlCli = writable({})
export const date = writable("")
export const cpuUsage = writable("")
export const memUsage = writable("")

export const clientSide = writable(5)
export const serverSide = writable(100)

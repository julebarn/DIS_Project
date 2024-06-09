import { writable } from 'svelte/store';
import { redirect } from '@sveltejs/kit';
import { goto } from "$app/navigation";;

const Storage = typeof localStorage !== "undefined" ? localStorage : {
    getItem: (key: string) => null,
    setItem: (key: string, value: string) => { },
    removeItem: (key: string) => { }
}


export const userID = writable<string | null>(Storage.getItem("userID"))
userID.subscribe(v => {
    if (v != null) {
        Storage.setItem("userID", v)
    } else {
        Storage.removeItem("userID")
    }
})




let isLoggin = false
userID.subscribe(v => isLoggin = v != null)

userID.subscribe(v => console.log("userID", v))

export function refreshTokens() {
    if (process.browser) {
        fetch("/api/auth/refresh", {
            credentials: "include",
        })
            .then(r => r.json())
            .then(v => {
                console.log("refresh");
                userID.set(v.userid)
                console.log(v);

                if (v.auth) {
                    setTimeout(refreshTokens, 20 * 60000)
                }
            })
    }
}


export function login(username: string, password: string) {
    if (process.browser) {
        fetch("/api/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ "username": username, "password": password })
        })
            .then(r => r.json())
            .then(v => {
                console.log(v);

                userID.set(v.userid)
                goto("/")
            })
    }


    setTimeout(refreshTokens, 20 * 60000)
}

export function logout() {

    if (process.browser) {
        fetch("/api/auth/logout")
    }

    userID.set(null)
}

export function register(username: string, password: string) {
    if (process.browser) {
        fetch("/api/auth/register", {
            method: "POST",
            body: JSON.stringify({ "username": username, "password": password })
        })
            .then(r => r.json())
            .then(v => {
                console.log(v);

                userID.set(v.userid)
                goto("/")
            })
        setTimeout(refreshTokens, 20 * 60000)
    }

}

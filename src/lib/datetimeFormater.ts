

export function datetimeFormater(datetime: string): string {
    const time = new Date(datetime)

    return `${time.getHours()}:${time.getMinutes()} - ${time.getDate()}/${time.getMonth()}/${time.getFullYear()}`
}
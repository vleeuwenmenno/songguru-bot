import { CircularProgress } from "@material-ui/core"
import React from "react"
import ResourceState from "../features/types"

interface ResourceProps<T> {
    selector: ResourceState<T>
    onError: React.ReactNode
    children: React.ReactNode
}

export function Resource<T>({
    selector,
    onError,
    children,
}: ResourceProps<T>): JSX.Element {
    if (!selector.loading && selector.error) {
        return <>{onError}</>
    }

    return <>
        {selector.loading && <CircularProgress />}
        {!selector.loading && selector.data && children}
        {!selector.loading && !selector.data && (
            <p>Provided data is undefined or false</p>
        )}
    </>

}

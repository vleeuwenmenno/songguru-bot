
interface ResourceState<T> {
    loading: boolean,
    data: T,
    error: string,
}

export default ResourceState
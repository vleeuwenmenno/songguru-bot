import { useDispatch } from "react-redux";
import { doLogout } from "../features/authentication/authenticationSlice";
import { Button } from "react-bootstrap"
import { AppDispatch } from "../store";


function LogoutButton(): JSX.Element {
    const dispatch = useDispatch<AppDispatch>()

    function logoutButtonClick(event: any): void {
        dispatch(doLogout())
        window.location.href = '/';
    }

    return <Button className="float-right" onClick={logoutButtonClick} variant="secondary">
        Logout
    </Button>
}

export default LogoutButton
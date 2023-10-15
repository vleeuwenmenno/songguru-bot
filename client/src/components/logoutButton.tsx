import { useDispatch } from "react-redux";
import { doLogout } from "../features/authentication/authenticationSlice";
import { Button } from "@material-ui/core";
import { AppDispatch } from "../store";


function LogoutButton(): JSX.Element {
    const dispatch = useDispatch<AppDispatch>()

    function logoutButtonClick(event: any): void {
        dispatch(doLogout())
        window.location.href = '/';
    }

    return <Button onClick={logoutButtonClick} variant="contained">
        Logout
    </Button>
}

export default LogoutButton
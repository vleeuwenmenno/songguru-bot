import { Button } from "@material-ui/core"

function LoginButton(): JSX.Element {
    function loginButtonClick(event: any): void {
        window.location.href = "http://localhost:8081/api/auth"
    }

    return <Button onClick={loginButtonClick} variant="contained">
        Login with Discord
    </Button>
}

export default LoginButton
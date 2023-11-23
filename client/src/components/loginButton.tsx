import { Button } from "react-bootstrap";
import { FaDiscord } from "react-icons/fa";
import './LoginButton.scss';

function LoginButton(): JSX.Element {
    function loginButtonClick(event: any): void {
        window.location.href = "http://localhost:8081/api/auth";
    }

    return <Button variant="primary" onClick={loginButtonClick} className="discord-btn">
        <FaDiscord size={24} className="discord-icon" />
        <span className="discord-text">Login with Discord</span>
    </Button>
}

export default LoginButton;

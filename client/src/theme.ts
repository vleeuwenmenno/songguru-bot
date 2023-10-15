import {  createTheme } from '@material-ui/core/styles';

const theme = createTheme({
  palette: {
    type: 'dark',
    background: {
      default: '#36393F',
      paper: '#2C2F33',
    },
    primary: {
      main: '#7289DA',
    },
    secondary: {
      main: '#FFFFFF',
    },
  },
});

export default theme;

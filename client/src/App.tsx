import { FC, useEffect } from 'react'
import { Button, Container, Card, CardContent, CardActions, Typography, ThemeProvider, createMuiTheme, ButtonGroup, Box, Tabs, Tab, AppBar } from '@material-ui/core'
import { fetchWhoAmI, doLogout } from './features/authentication/authenticationSlice'
import { useDispatch, useSelector } from 'react-redux'
import { AppDispatch, RootState } from './store'
import LoginButton from './components/loginButton'
import { fetchPreferences } from './features/preferences/preferencesSlice'
import './App.scss'

const App: FC = () => {
  const dispatch = useDispatch<AppDispatch>()
  const auth = useSelector((state: RootState) => state.auth)

  useEffect(() => {
    dispatch(fetchWhoAmI())
  }, [dispatch])

  useEffect(() => {
    if (auth.isAuthenticated) {
      dispatch(fetchPreferences())
    }
  }, [dispatch, auth.isAuthenticated])

  return (
    <Container maxWidth="sm">
      <Card>
        <CardContent>
          <Typography variant="h5" component="h2" align="center">
            {auth.isAuthenticated ? 'Welcome back, ' + auth.data.global_name : ''}
          </Typography>
        </CardContent>
        {auth.isAuthenticated ? (
          <CardActions style={{ justifyContent: 'right', marginRight: 8 }}>
            {auth.isAuthenticated ? (
              <Button color="primary" onClick={() => dispatch(doLogout())}>
                Logout
              </Button>
            ) : (
              <LoginButton />
            )}
          </CardActions>) : (
          <CardActions style={{ justifyContent: 'center' }}>
            {auth.isAuthenticated ? (
              <Button color="primary" onClick={() => dispatch(doLogout())}>
                Logout
              </Button>
            ) : (
              <LoginButton />
            )}
          </CardActions>)}
      </Card>
    </Container>
  )
}

export default App

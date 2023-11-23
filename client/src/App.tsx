import { FC, useEffect } from 'react'
import { fetchWhoAmI, doLogout } from './features/authentication/authenticationSlice'
import { useDispatch, useSelector } from 'react-redux'
import { AppDispatch, RootState } from './store'
import LoginButton from './components/loginButton'
import { fetchPreferences } from './features/preferences/preferencesSlice'
import './App.scss'
import { Accordion, Spinner, Tabs, Tab, Card, Container, CardBody, CardFooter } from 'react-bootstrap'
import LogoutButton from './components/logoutButton'

const App: FC = () => {
  const dispatch = useDispatch<AppDispatch>()
  const preferences = useSelector((state: RootState) => state.preferences)
  const auth = useSelector((state: RootState) => state.auth)

  useEffect(() => {
    dispatch(fetchWhoAmI())
  }, [dispatch])

  useEffect(() => {
    if (auth.isAuthenticated) {
      dispatch(fetchPreferences())
    }
  }, [dispatch, auth.isAuthenticated])

  if (auth.loading || preferences.loading) {
    return (
      <Container className='h-screen flex items-center justify-center'>
        <Spinner />
      </Container>
    )
  }

  return (
    <Container className='h-screen flex items-center justify-center'>
      <Card className='w'>
        <CardBody>
          <p className="text-3xl">
            {auth.isAuthenticated ? 'Welcome back, ' + auth.data.global_name : 'Login to continue'}
          </p>
          {!auth.isAuthenticated ? <p className="text-center"><LoginButton /></p> : <></>}
          {auth.isAuthenticated ?
            <Tabs
              defaultActiveKey="guilds"
              id="uncontrolled-tab-example"
              className="mb-3"
            >
              <Tab eventKey="guilds" title="Guilds">
                <Accordion defaultActiveKey="0">
                  {preferences.data.map((guild, index) => (
                    <Accordion.Item eventKey={index.toString()}>
                      <Accordion.Header>{guild.guild.name}</Accordion.Header>
                      <Accordion.Body>
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
                        eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
                        minim veniam, quis nostrud exercitation ullamco laboris nisi ut
                        aliquip ex ea commodo consequat. Duis aute irure dolor in
                        reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
                        pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
                        culpa qui officia deserunt mollit anim id est laborum.
                      </Accordion.Body>
                    </Accordion.Item>
                  ))}
                </Accordion>
              </Tab>
              <Tab eventKey="history" title="History">
                History
              </Tab>
              <Tab eventKey="defaults" title="Defaults">
                Tab content for Contact
              </Tab>
            </Tabs> : <></>}
        </CardBody>
        {auth.isAuthenticated ?
          <CardFooter>
            <LogoutButton />
          </CardFooter> : <></>}
      </Card>
    </Container>
  )
}

export default App

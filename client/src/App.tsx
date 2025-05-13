import { Stack, Container } from "@chakra-ui/react"
import TodoForm from "./components/TodoForm"
import Navbar from "./components/Navbar"
import TodoList from "./components/TodoList"

export const BASE_URL = "http://localhost:4000/api/"
function App() {

  return (
    <Stack h={"100vh"} >
      <Navbar></Navbar>
      <Container maxW={"700px"}>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  )
}

export default App

import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";
import TodoItem from "./TodoItem";
import { useQuery } from "@tanstack/react-query";

export type Task = {
    _id: number;
    body: string;
    completed: boolean;
}
const TodoList = () => {

    const { data: tasks, isLoading } = useQuery<Task[]>({
        queryKey: ["tasks"],
        queryFn: async () => {
            try {
                const response = await fetch("http://localhost:4000/api/tasks")
                const data = await response.json();

                if (!response.ok) {
                    throw new Error(data.message || "Something went wrong");
                }
                return data || [];
            } catch (error) {
                console.error("Error fetching tasks:", error);
                return [];
            }
        }
    })
    return (
        <>
            <Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}>
                Today's Tasks
            </Text>
            {isLoading && (
                <Flex justifyContent={"center"} my={4}>
                    <Spinner size={"xl"} />
                </Flex>
            )}
            {!isLoading && tasks?.length === 0 && (
                <Stack alignItems={"center"} gap='3'>
                    <Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
                        All tasks completed! 🤞
                    </Text>
                </Stack>
            )}
            <Stack gap={3}>
                {tasks?.map((task, idx) => (
                    <TodoItem key={idx} todo={task} />
                ))}
            </Stack>
        </>
    );
};
export default TodoList;
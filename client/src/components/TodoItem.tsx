import { Badge, Box, Flex, Spinner, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import type { Task } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BASE_URL } from "@/App";

const TodoItem = ({ todo }: { todo: Task }) => {
	const queryClient = useQueryClient();

	const { mutate: updateTask, isPending: isUpdating } = useMutation({
		mutationKey: ["updateTask"],
		mutationFn: async () => {
			try {
				const response = await fetch(`${BASE_URL}tasks/${todo._id}`, {
					method: "PATCH",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({ completed: !todo.completed }),
				});
				const data = await response.json();

				if (!response.ok) {
					throw new Error(data.message || "Something went wrong");
				}
				return data;
			} catch (error) {
				console.error("Error updating task:", error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["tasks"] });
		}
	})

	const { mutate: deleteTask, isPending: isDeleting } = useMutation({
		mutationKey: ["deleteTask"],
		mutationFn: async () => {
			try {
				const response = await fetch(`${BASE_URL}tasks/${todo._id}`, {
					method: "DELETE",
					headers: {
						"Content-Type": "application/json",
					},
				});
				const data = await response.json();

				if (!response.ok) {
					throw new Error(data.message || "Something went wrong");
				}
				return data;
			} catch (error) {
				console.error("Error deleting task:", error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["tasks"] });
		}
	})

	return (
		<Flex gap={2} alignItems={"center"}>
			<Flex
				flex={1}
				alignItems={"center"}
				border={"1px"}
				borderColor={"gray.900"}
				p={2}
				borderRadius={"lg"}
				justifyContent={"space-between"}
			>
				<Text
					color={todo.completed ? "green.200" : "yellow.100"}
					textDecoration={todo.completed ? "line-through" : "none"}
				>
					{todo.body}
				</Text>
				{todo.completed && (
					<Badge ml='1' colorScheme='green'>
						Done
					</Badge>
				)}
				{!todo.completed && (
					<Badge ml='1' colorScheme='yellow'>
						In Progress
					</Badge>
				)}
			</Flex>
			<Flex gap={2} alignItems={"center"}>
				<Box color={"green.500"} cursor={"pointer"} onClick={updateTask}>
					{!isUpdating && <FaCheckCircle size={20} />}
					{isUpdating && <Spinner size={"sm"} />}

				</Box>
				<Box color={"red.500"} cursor={"pointer"} onClick={deleteTask}>
					{!isDeleting && <MdDelete size={20} />}
					{isDeleting && <Spinner size={"sm"} />}
				</Box>
			</Flex>
		</Flex>
	);
};
export default TodoItem;
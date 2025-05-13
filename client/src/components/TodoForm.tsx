import { BASE_URL } from "@/App";
import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";

const TodoForm = () => {
	const [newTask, setNewTask] = useState("");

	const queryClient = useQueryClient();

	const { mutate: createTask, isPending: isCreating } = useMutation({
		mutationKey: ["createTask"],
		mutationFn: async (e: React.FormEvent) => {
			e.preventDefault();
			try {
				const response = await fetch(BASE_URL + "tasks", {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
					},
					body: JSON.stringify({ body: newTask }),
				})
				const data = await response.json();
				if (!response.ok) {
					throw new Error(data.message || "Something went wrong");
				}
				setNewTask("");
				return data;
			} catch (error: any) {
				console.error("Error creating task:", error);
			}
		},
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["tasks"] });
		},
		onError: (error) => {
			alert(error.message);
		}
	})

	return (
		<form onSubmit={createTask}>
			<Flex gap={2}>
				<Input
					type='text'
					value={newTask}
					onChange={(e) => setNewTask(e.target.value)}
					ref={(input) => input && input.focus()}
				/>
				<Button
					mx={2}
					type='submit'
					_active={{
						transform: "scale(.97)",
					}}
				>
					{isCreating ? <Spinner size={"xs"} /> : <IoMdAdd size={30} />}
				</Button>
			</Flex>
		</form>
	);
};
export default TodoForm;
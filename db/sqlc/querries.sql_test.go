package db

import (
	"context"
	f "github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username: f.Username(),
		Password: f.Password(true, true, true, true, false, 10),
		Email:    f.Email(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)

	return user
}

func createRandomGroup(t *testing.T) Groups {
	arg := CreateGroupParams{
		Name: f.Word(),
	}
	group, err := testStore.CreateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group)
	require.Equal(t, arg.Name, group.Name)

	return group

}

func createRandomTask(t *testing.T) Tasks {
	description := f.Sentence(20)
	dueDate := f.Date()

	arg := CreateTaskParams{
		Title:       f.Word(),
		Description: &description,
		DueDate: pgtype.Timestamp{
			Time:  dueDate,
			Valid: true,
		},
		Status: "todo",
	}
	task, err := testStore.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.WithinDuration(t, dueDate, task.DueDate.Time, 10*time.Second)
	require.Equal(t, arg.Status, task.Status)

	return task

}

func createRandomSubtask(t *testing.T, taskID int32) Subtasks {
	description := f.Sentence(20)
	dueDate := f.Date()

	arg := CreateSubtaskParams{
		Title:       f.Word(),
		Description: &description,
		DueDate: pgtype.Timestamp{
			Time:  dueDate,
			Valid: true,
		},
		Status: "todo",
		TaskID: &taskID,
	}
	subtask, err := testStore.CreateSubtask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, subtask)
	require.Equal(t, arg.Title, subtask.Title)
	require.Equal(t, arg.Description, subtask.Description)
	require.WithinDuration(t, dueDate, subtask.DueDate.Time, 10*time.Second)
	require.Equal(t, arg.Status, subtask.Status)

	return subtask

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testStore.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)
	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	newEmail := f.Email()
	newPassword := f.Password(true, true, true, true, false, 10)

	arg := UpdateUserParams{
		Username: &user1.Username,
		Password: &newPassword,
		Email:    &newEmail,
	}
	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, arg.Username, &user2.Username)
	require.Equal(t, arg.Password, &user2.Password)
	require.Equal(t, arg.Email, &user2.Email)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.Empty(t, user2)
}

func TestCreateGroup(t *testing.T) {
	arg := CreateGroupParams{
		Name: f.Word(),
	}
	group, err := testStore.CreateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group)
	require.Equal(t, arg.Name, group.Name)
}

func TestGetGroup(t *testing.T) {
	group1 := createRandomGroup(t)
	group2, err := testStore.GetGroup(context.Background(), group1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, group2)
	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, group1.Name, group2.Name)
}

func TestListGroups(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomGroup(t)
	}
	args := ListGroupsParams{
		Limit:  10,
		Offset: 0,
	}

	groups, err := testStore.ListGroups(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, groups, 10)
	for _, group := range groups {
		require.NotEmpty(t, group)
	}
}

func TestUpdateGroup(t *testing.T) {
	group1 := createRandomGroup(t)
	newName := f.Word()
	arg := UpdateGroupParams{
		ID:   group1.ID,
		Name: &newName,
	}
	group2, err := testStore.UpdateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group2)
	require.Equal(t, arg.Name, &group2.Name)
}

func TestDeleteGroup(t *testing.T) {
	group1 := createRandomGroup(t)
	err := testStore.DeleteGroup(context.Background(), group1.ID)
	require.NoError(t, err)
	group2, err := testStore.GetGroup(context.Background(), group1.ID)
	require.Error(t, err)
	require.Empty(t, group2)
}

func TestAddUserToGroup(t *testing.T) {
	user := createRandomUser(t)
	group := createRandomGroup(t)
	arg := AddUserToGroupParams{
		UserID:  user.ID,
		GroupID: group.ID,
	}
	err := testStore.AddUserToGroup(context.Background(), arg)
	require.NoError(t, err)
}

func TestRemoveUserFromGroup(t *testing.T) {
	user := createRandomUser(t)
	group := createRandomGroup(t)
	arg := AddUserToGroupParams{
		UserID:  user.ID,
		GroupID: group.ID,
	}
	err := testStore.AddUserToGroup(context.Background(), arg)
	require.NoError(t, err)

	arg2 := RemoveUserFromGroupParams{
		UserID:  user.ID,
		GroupID: group.ID,
	}
	err = testStore.RemoveUserFromGroup(context.Background(), arg2)
	require.NoError(t, err)

}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testStore.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.Description, task2.Description)
	require.Equal(t, task1.DueDate, task2.DueDate)
	require.Equal(t, task1.Status, task2.Status)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTask(t)
	}

	args := ListTasksParams{
		Limit:  5,
		Offset: 5,
	}

	tasks, err := testStore.ListTasks(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, tasks, 5)
	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}

func TestUpdateTask(t *testing.T) {
	task1 := createRandomTask(t)
	newDescription := f.Sentence(20)
	newDueDate := f.Date()
	newStatus := "done"

	arg := UpdateTaskParams{
		ID:          task1.ID,
		Title:       &task1.Title,
		Description: &newDescription,
		DueDate: pgtype.Timestamp{
			Time:  newDueDate,
			Valid: true,
		},
		Status: &newStatus,
	}
	task2, err := testStore.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)
	require.Equal(t, arg.Title, &task2.Title)
	require.Equal(t, arg.Description, task2.Description)
	require.WithinDuration(t, newDueDate, task2.DueDate.Time, 10*time.Second)
	require.Equal(t, arg.Status, &task2.Status)
}

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)
	err := testStore.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)
	task2, err := testStore.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.Empty(t, task2)
}

func TestCreateSubtask(t *testing.T) {
	task := createRandomTask(t)
	createRandomSubtask(t, task.ID)
}

func TestGetSubtask(t *testing.T) {
	task := createRandomTask(t)
	subtask1 := createRandomSubtask(t, task.ID)
	subtask2, err := testStore.GetSubtask(context.Background(), subtask1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, subtask2)
	require.Equal(t, subtask1.ID, subtask2.ID)
	require.Equal(t, subtask1.Title, subtask2.Title)
	require.Equal(t, subtask1.Description, subtask2.Description)
	require.Equal(t, subtask1.DueDate, subtask2.DueDate)
	require.Equal(t, subtask1.Status, subtask2.Status)
}

func TestListTaskSubtasks(t *testing.T) {
	task := createRandomTask(t)
	for i := 0; i < 10; i++ {
		createRandomSubtask(t, task.ID)
	}

	subtasks, err := testStore.ListTaskSubtasks(context.Background(), &task.ID)
	require.NoError(t, err)
	require.Len(t, subtasks, 10)
	for _, subtask := range subtasks {
		require.NotEmpty(t, subtask)
		require.Equal(t, task.ID, *subtask.TaskID)
	}
}

func TestUpdateSubtask(t *testing.T) {
	task := createRandomTask(t)
	subtask1 := createRandomSubtask(t, task.ID)
	newDescription := f.Sentence(20)
	newDueDate := f.Date()
	newStatus := "done"

	arg := UpdateSubtaskParams{
		ID:          subtask1.ID,
		Title:       &subtask1.Title,
		Description: &newDescription,
		DueDate: pgtype.Timestamp{
			Time:  newDueDate,
			Valid: true,
		},
		Status: &newStatus,
	}
	subtask2, err := testStore.UpdateSubtask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, subtask2)
	require.Equal(t, arg.Title, &subtask2.Title)
	require.Equal(t, arg.Description, subtask2.Description)
	require.WithinDuration(t, newDueDate, subtask2.DueDate.Time, 10*time.Second)
	require.Equal(t, arg.Status, &subtask2.Status)
}

func TestDeleteSubtask(t *testing.T) {
	task := createRandomTask(t)
	subtask1 := createRandomSubtask(t, task.ID)
	err := testStore.DeleteSubtask(context.Background(), subtask1.ID)
	require.NoError(t, err)
	subtask2, err := testStore.GetSubtask(context.Background(), subtask1.ID)
	require.Error(t, err)
	require.Empty(t, subtask2)
}

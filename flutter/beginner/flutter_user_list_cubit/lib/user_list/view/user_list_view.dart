part of "user_list_page.dart";

class UserListView extends StatelessWidget {
  const UserListView({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: BlocBuilder<UserListCubit, UserListState>(
        builder: (context, state) {
          if (state is UserListSuccess) {
            // Success
            return Scaffold(
              body: ListView.builder(
                itemCount: state.users.length,
                itemBuilder: (context, index) {
                  return ListTile(
                    title: Text(state.users[index].name),
                    subtitle: Text(state.users[index].email),
                  );
                },
              ),
              floatingActionButton: FloatingActionButton(
                child: const Icon(Icons.refresh),
                onPressed: () => context.read<UserListCubit>().fetchUser(),
              ),
            );
          } else if (state is UserListError) {
            // Error
            return Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(state.errorMessage),
                  ElevatedButton(
                    child: const Text("Refresh"),
                    onPressed: () => context.read<UserListCubit>().fetchUser(),
                  ),
                ],
              ),
            );
          } else if (state is UserListLoading) {
            // Loading
            return const Center(
              child: CircularProgressIndicator(),
            );
          } else {
            // Initial / First time
            return Center(
              child: ElevatedButton(
                child: const Text("Refresh"),
                onPressed: () => context.read<UserListCubit>().fetchUser(),
              ),
            );
          }
        },
      ),
    );
  }
}

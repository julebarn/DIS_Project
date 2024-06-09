GITHUB LINK: https://github.com/julebarn/DIS_Project

# DIS_Project

# Compiling and Running through Docker

If you want to locally host the server and website through Docker, 
you can do so by simply running

```bash
docker-compose up --build
```

When the container is running, you can access the website at ``http://localhost:8080/``.

# Teardown

If you wish to delete the containers then please run 

```bash
docker compose rm
```

inside the project's root directory.

If you wish to remove associated volumes, then please run

```bash
docker volume rm dis_project_db-data
```

If you wish to remove associated images, then please run

```bash
docker image rm dis_project-server
```

# Regexes
Regexes used can be found in ``src/routes/register/+page.svelte``.

# Sample Database 
The sample database can be found in ``init-data.sql``. The sample database should be automatically inserted when running the application. 
The sample database contains two users;
- User1: Username = "Test1" Password = "Test1234%"
- User2: Username = "Test2" Password = "Test1234%"
With 5 events and 2 clubs. 


# Known Issues
- Sometimes when creating a new event, the event is not shown at the front-end, but we do believe that it is inserted into the database. This may be caused by the front-end only showing future events.
- When adding a user as manager in a club, you may get an issue that the addition of the member violates foreign key contraint. This is caused by the frontend setting the default user as ``null``, 
  which means you have to select a user in the list before you add it, even if it seems like it is already selected.

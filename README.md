# stupid-simple-blog

A stupid simple blog service written in Scala, Akka, and PSQL.

# SQL structure

### posts

| id  | user_id *(fk users)* | title | created_at |
|-----|----------------------|-------|------------|

### post_contents

| id  | post_id *(fk posts)* | content |
|-----|----------------------|---------|

### users

| id  | f_name | l_name |
|-----|--------|--------|

### comments

| id  | user_id *(fk users)* | post_id *(fk posts)* | content | created_at |
|-----|----------------------|----------------------|---------|------------|

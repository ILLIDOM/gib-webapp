import sqlite3
from sqlite3 import Error

import sqlite3
from sqlite3 import Error


def create_sqlite_db(db_file):
    conn = None
    try:
        conn = sqlite3.connect(db_file)
        c = conn.cursor()

        # create table users
        c.execute('''
        CREATE TABLE users 
        (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            firstname TEXT,
            lastname TEXT,
            fullname TEXT,
            email TEXT,
            password TEXT
        )
        ''')

        c.execute('''
        INSERT INTO users VALUES
        (
            NULL,
            'Dominique',
            'Illi',
            'Dominique Illi',
            'test@example.com',
            'password'
        ),
        (
            NULL,
            'Fabian',
            'Thurnheer',
            'Fabian Thurnheer',
            'test@change.com',
            'password'
        )
        ''')

        # create roles table
        c.execute('''
        CREATE TABLE roles 
        (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            name TEXT
        )
        ''')

        c.execute('''
        INSERT INTO roles VALUES
            (NULL,  'Admin'),
            (NULL, 'Finanzen')
        ''')


        # create table userrole
        c.execute('''
        CREATE TABLE userrole
        (
            id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
            userid INTEGER,
            roleid INTEGER,
            FOREIGN KEY(userid) REFERENCES users(userid),
            FOREIGN KEY(roleid) REFERENCES roles(roleid)
        )
        ''')

        c.execute('''
        INSERT INTO userrole VALUES
            (NULL, 1, 1),
            (NULL, 2, 2)
        ''')

        conn.commit()
    except Error as e:
        print(e)
    finally:
        if conn:
            conn.close()


if __name__ == '__main__':
    create_sqlite_db(r"database.db")
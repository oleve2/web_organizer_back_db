
from datetime import *
import sqlite3

"""
before launching script: 
- setup final sqilte db name
- change dates in "make_dates" function to whatever you need
"""
# ------------------------------------
def make_tables(conn):
  sql_01_01_posts = """
  CREATE TABLE "posts" (
    "id"    INTEGER,
    "title" TEXT,
    "text"  TEXT,
    "theme" TEXT,
    "part"  TEXT,
    "tags_list" TEXT,
    PRIMARY KEY("id" AUTOINCREMENT)
  );
  """

  sql_01_02_activ_log = """
  CREATE TABLE "activ_log" (
    "id"    INTEGER,
    "activ_name_id" INTEGER NOT NULL,
    "activ_norm_id" INTEGER NOT NULL,
    "activ_date"    TEXT NOT NULL,
    "activ_value"   INTEGER NOT NULL,
    PRIMARY KEY("id" AUTOINCREMENT)
  );
  """

  sql_01_03_activ_names = """
  CREATE TABLE "activ_names" (
    "id"    INTEGER,
    "name"  TEXT UNIQUE,
    "date_start"    TEXT,
    "date_end"  TEXT,
    "norm_id"   INTEGER,
    PRIMARY KEY("id" AUTOINCREMENT)
  );
  """

  sql_01_04_activ_normative = """
  CREATE TABLE "activ_normative" (
    "id"	INTEGER,
    "name"	TEXT,
    "norm_period"	TEXT,
    "norm_value"	INTEGER,
    "norm_measure"	TEXT,
    PRIMARY KEY("id" AUTOINCREMENT)
  );
  """

  sql_01_05_dates = """
  CREATE TABLE "dates" (
    "date"	TEXT UNIQUE,
    PRIMARY KEY("date")
  );
  """

  cur1 = conn.cursor()

  sql_list = [
      sql_01_01_posts,
      sql_01_02_activ_log,
      sql_01_03_activ_names,
      sql_01_04_activ_normative,
      sql_01_05_dates
  ]

  for elem in sql_list:
      print(elem)
      cur1.execute(elem)
      print(" ==> done \n -------------------------\n")

  conn.commit()

# ------------------------------------
def make_dates(conn):
  total_dates = []
  date_from = date.fromisoformat('2023-01-01')
  date_to = date.fromisoformat('2025-01-01')

  curr_date = date_from
  while (curr_date <= date_to):
      d = ( str(curr_date), ) #{"date": str(curr_date)}
      #print(d)
      total_dates.append(d)
      curr_date += timedelta(days=1)

  # вставка данных
  def insert_date(conn, dataObj):
      for elem in dataObj:
          #print(elem)
          #input()
          sql = """insert into dates(date) VALUES(?)"""
          cur = conn.cursor()
          cur.execute(sql, elem ) 
      print('done')

  #
  insert_date(conn, total_dates)

  conn.commit()


# ------------------------------------
def insert_dummy_posts(conn):
  sql_insert_01 = """
  insert into posts (title, text, theme, part, tags_list) values
  ('A1 title 1','some random text A1','P1 Theme 1','PART 1', '[]'),
  ('A1 title 2','some random text A2','P1 Theme 2','PART 1', '[]'),
  ('A1 title 1','some random text A3','P2 Theme 1','PART 2', '[]')
  """
  cur1 = conn.cursor()
  cur1.execute(sql_insert_01)
  conn.commit()
  print('posts dummy data inserted')

   


# ====================================
if __name__ == '__main__':
  conn = sqlite3.connect('./prd.db')
  # tables
  make_tables(conn)
  # dates
  make_dates(conn)
  # dummy posts data
  insert_dummy_posts(conn)
  
  conn.close()


#!/usr/local/bin/python3
import json
import sys
import os

jsd=sys.stdin.read()

d = json.loads(jsd)

def delete_file_if_exists(filename):
  if os.path.exists(filename):
    os.remove(filename)

def save_data_to_file(filename,util,data):
  ret=[str(util),str(data)]
  s="\t".join(ret)
  print(s,file=filename)

with open("randread.dat","a") as randread,open("randwrite.dat","a") as randwrite,open("seqread.dat","a") as seqread,open("seqwrite.dat","a")as seqwrite:
  for oneret in d:
    for data in oneret:
      #time='"'+data["time"]+'"'
      util=data["disk_util"][0]["util"]
      jobname=data["jobs"][0]["jobname"]
      want=''
      if jobname=='4k_randwrite':
        want=data["jobs"][0]["write"]["iops"]
        save_data_to_file(randwrite,util,want)
      elif jobname=='4k_randread':
        want=data["jobs"][0]["read"]["iops"]
        save_data_to_file(randread,util,want)
      elif jobname=='1024k_write':
        want=data["jobs"][0]["write"]["bw"]
        save_data_to_file(seqwrite,util,want)
      elif jobname=='1024k_read':
        want=data["jobs"][0]["read"]["bw"]
        save_data_to_file(seqread,util,want)
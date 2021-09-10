import subprocess

while True:
    try:
        res = subprocess.check_output(
            ['python', 'raritygems.py'],
            universal_newlines=True, stderr=subprocess.STDOUT)
    except Exception as e:
        print(e.stdout[:-1])
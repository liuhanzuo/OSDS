import os
import glob

def clean_exe_files(directory):
    exe_files = glob.glob(os.path.join(directory, '*.exe'))
    for exe_file in exe_files:
        try:
            os.remove(exe_file)
            print(f"Removed: {exe_file}")
        except Exception as e:
            print(f"Error removing {exe_file}: {e}")

def clean_txt_files(directory):
    exe_files = glob.glob(os.path.join(directory, '*.txt'))
    for exe_file in exe_files:
        try:
            os.remove(exe_file)
            print(f"Removed: {exe_file}")
        except Exception as e:
            print(f"Error removing {exe_file}: {e}")

def clean_png_files(directory):
    exe_files = glob.glob(os.path.join(directory, '*.png'))
    for exe_file in exe_files:
        try:
            os.remove(exe_file)
            print(f"Removed: {exe_file}")
        except Exception as e:
            print(f"Error removing {exe_file}: {e}")

if __name__ == "__main__":
    directory = './'
    clean_exe_files(directory)
    directory = './data'
    clean_txt_files(directory)
    directory = './fig'
    clean_png_files(directory)


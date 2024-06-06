
import json
import os


def parse_log_file(log_file_path, output_dir):
  """
  Parses a log file for lines containing "Response=<json>" and extracts JSON content, 
  writing each to a separate file.

  Args:
    log_file_path: Path to the log file.
    output_dir: Directory to write the extracted JSON files.
  """
  with open(log_file_path, 'r') as log_file:
    file_name = 1
    for line in log_file:
      # Check if line contains "Response=" followed by "<json>" pattern
      if "Response=" in line:
        # Extract the JSON content between "<" and ">"
        start_index = line.find("Response=") + 9
        json_content = line[start_index:].strip()
        print(json_content)
        try:
          # Try to parse the extracted content as JSON
          json_data = json.loads(json_content)
          # Create a filename with sequential numbering
          output_file_name = f"{output_dir}/ndfc_intf_vlan_{file_name}.json"
          with open(output_file_name, 'w') as output_file:
            json.dump(json_data, output_file)
          file_name += 1
        except json.JSONDecodeError:
          # Handle lines that don't have valid JSON after "Response=<json>" (optional)
          pass




# Replace with your actual paths
log_file_path = "/tmp/terraform-acceptance-tests.log"
output_dir = "./interface_vlan"
os.mkdir(output_dir)

parse_log_file(log_file_path, output_dir)

print(f"Finished parsing log file. Extracted JSON files are in {output_dir}")
